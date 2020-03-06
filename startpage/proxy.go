package startpage

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"bakku.org/sherlock"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

const baseURL = "https://startpage.com/do/search"
const resultsXPath = "//div[@class=\"w-gl__result\"]"
const linkXPath = ".//a[@class=\"w-gl__result-title\"]"
const descriptionXPath = ".//p[@class=\"w-gl__description\"]"

// Proxy can fetch search results from startpage.com for a given query.
type Proxy struct {
	httpClient *http.Client
}

// FetchResults fetches all search results for a given query.
// Error can occur in various cases so it should be checked!
func (p *Proxy) FetchResults(query string) ([]sherlock.SearchResult, error) {
	formData := url.Values{
		"query": {query},
		"page":  {"1"},
		"cat":   {"web"},
		"cmd":   {"process_search"},
	}

	resp, err := p.httpClient.PostForm(baseURL, formData)
	if err != nil {
		return nil, fmt.Errorf("error during request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("got status " + resp.Status + " from startpage")
	}

	return extractResultsFromHTML(resp.Body)
}

func extractResultsFromHTML(html io.Reader) ([]sherlock.SearchResult, error) {
	parsedDoc, err := htmlquery.Parse(html)
	if err != nil {
		return nil, fmt.Errorf("could not parse startpage html: %v", err)
	}

	// ignore error since it is a valid xpath expression
	nodes, _ := htmlquery.QueryAll(parsedDoc, resultsXPath)

	results := []sherlock.SearchResult{}

	for _, node := range nodes {
		// ignore errors since it is a valid xpath expression
		link, _ := htmlquery.Query(node, linkXPath)
		description, _ := htmlquery.Query(node, descriptionXPath)

		if link == nil || description == nil {
			continue
		}

		result := sherlock.SearchResult{}

		result.Title = extractText(link)
		result.Description = extractText(description)
		result.Link = extractHref(link)

		results = append(results, result)
	}

	return results, nil
}

func extractText(node *html.Node) string {
	innerText := htmlquery.InnerText(node)
	r, _ := regexp.Compile("(\\s{2,}|\u00a0|\n)")
	return r.ReplaceAllString(strings.TrimSpace(innerText), " ")
}

func extractHref(node *html.Node) string {
	attributes := node.Attr

	for _, attr := range attributes {
		if attr.Key == "href" {
			return attr.Val
		}
	}

	return ""
}

// NewProxy returns an instance of a startpage proxy scraper
func NewProxy(client *http.Client) *Proxy {
	return &Proxy{client}
}
