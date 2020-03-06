package startpage_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"bakku.org/sherlock/startpage"
)

type MockTransport struct {
	req  *http.Request
	resp *http.Response
	err  error
}

func (m MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	m.req = req
	return m.resp, m.err
}

func TestFetchResults_ShouldReturnCorrectResultsForValidHTML(t *testing.T) {
	bytes, err := ioutil.ReadFile("./test_files/valid.html")
	if err != nil {
		t.Fatalf("could not read \"test_files/valid.html\": %v", err)
	}

	mockTransport := &MockTransport{
		resp: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(string(bytes))),
			Header:     make(http.Header),
		},
		err: nil,
	}

	mockHttpClient := &http.Client{
		Transport: mockTransport,
	}

	proxy := startpage.NewProxy(mockHttpClient)
	results, err := proxy.FetchResults("hello")
	if err != nil {
		t.Fatalf("expected no error but got %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 results but got %v", len(results))
	}

	first := results[0]

	expectedTitle := "Adele - Hello - YouTube"
	if first.Title != expectedTitle {
		t.Fatalf("expected title to be %#v but was %#v", expectedTitle, first.Title)
	}

	expectedDescription := "22 Oct 2015 ... 'Hello' is taken from the new album, 25, out " +
		"November 20. http://adele.com Available now from iTunes http://smarturl.it/itunes25 Available now ..."
	if first.Description != expectedDescription {
		t.Fatalf("expected description to be %#v but was %#v", expectedDescription, first.Description)
	}

	expectedLink := "https://www.youtube.com/watch?v=YQHsXMglC9A"
	if first.Link != expectedLink {
		t.Fatalf("expected link to be %#v but was %#v", expectedLink, first.Link)
	}

	second := results[1]

	expectedTitle = "Adele - Hello (Live at the NRJ Awards) - YouTube"
	if second.Title != expectedTitle {
		t.Fatalf("expected title to be %#v but was %#v", expectedTitle, second.Title)
	}

	expectedDescription = "9 Nov 2015 ... 'Hello' is taken from the new album, 25, out November 20. " +
		"http://adele.com Available now from iTunes http://smarturl.it/itunes25 Available now ..."
	if second.Description != expectedDescription {
		t.Fatalf("expected description to be %#v but was %#v", expectedDescription, second.Description)
	}

	expectedLink = "https://www.youtube.com/watch?v=DfG6VKnjrVw"
	if second.Link != expectedLink {
		t.Fatalf("expected link to be %#v but was %#v", expectedLink, second.Link)
	}
}

func TestFetchResults_ShouldReturnAnErrorInCaseRequestReturnsError(t *testing.T) {
	mockTransport := &MockTransport{
		resp: nil,
		err:  errors.New("some error"),
	}

	mockHttpClient := &http.Client{
		Transport: mockTransport,
	}

	proxy := startpage.NewProxy(mockHttpClient)
	_, err := proxy.FetchResults("hello")
	if err == nil {
		t.Fatalf("expected an error but got none")
	}
}

func TestFetchResults_ShouldReturnAnErrorInCaseStatusNotOK(t *testing.T) {
	mockTransport := &MockTransport{
		resp: &http.Response{
			StatusCode: http.StatusBadRequest,
			Body:       ioutil.NopCloser(strings.NewReader("")),
			Header:     make(http.Header),
		},
		err: nil,
	}

	mockHttpClient := &http.Client{
		Transport: mockTransport,
	}

	proxy := startpage.NewProxy(mockHttpClient)
	_, err := proxy.FetchResults("hello")
	if err == nil {
		t.Fatalf("expected an error but got none")
	}
}

func TestFetchResults_ShouldIgnoreResultsWithoutFetchableLink(t *testing.T) {
	bytes, err := ioutil.ReadFile("./test_files/invalid_link.html")
	if err != nil {
		t.Fatalf("could not read \"test_files/invalid_link.html\": %v", err)
	}

	mockTransport := &MockTransport{
		resp: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(string(bytes))),
			Header:     make(http.Header),
		},
		err: nil,
	}

	mockHttpClient := &http.Client{
		Transport: mockTransport,
	}

	proxy := startpage.NewProxy(mockHttpClient)
	results, err := proxy.FetchResults("hello")
	if err != nil {
		fmt.Println(err)
		t.Fatalf("expected no error to occur during fetching")
	}

	if len(results) != 0 {
		t.Fatalf("expected no results")
	}
}

func TestFetchResults_ShouldIgnoreResultsWithoutFetchableDescription(t *testing.T) {
	bytes, err := ioutil.ReadFile("./test_files/invalid_description.html")
	if err != nil {
		t.Fatalf("could not read \"test_files/invalid_description.html\": %v", err)
	}

	mockTransport := &MockTransport{
		resp: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(string(bytes))),
			Header:     make(http.Header),
		},
		err: nil,
	}

	mockHttpClient := &http.Client{
		Transport: mockTransport,
	}

	proxy := startpage.NewProxy(mockHttpClient)
	results, err := proxy.FetchResults("hello")
	if err != nil {
		fmt.Println(err)
		t.Fatalf("expected no error to occur during fetching")
	}

	if len(results) != 0 {
		t.Fatalf("expected no results")
	}
}

func TestFetchResults_ShouldReturnNoResultsForNoResults(t *testing.T) {
	bytes, err := ioutil.ReadFile("./test_files/no_results.html")
	if err != nil {
		t.Fatalf("could not read \"test_files/no_results.html\": %v", err)
	}

	mockTransport := &MockTransport{
		resp: &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(strings.NewReader(string(bytes))),
			Header:     make(http.Header),
		},
		err: nil,
	}

	mockHttpClient := &http.Client{
		Transport: mockTransport,
	}

	proxy := startpage.NewProxy(mockHttpClient)
	results, err := proxy.FetchResults("hello")
	if err != nil {
		fmt.Println(err)
		t.Fatalf("expected no error to occur during fetching")
	}

	if len(results) != 0 {
		t.Fatalf("expected no results")
	}
}
