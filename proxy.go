package sherlock

type SearchEngineProxy interface {
	FetchResults(query string) ([]SearchResult, error)
}
