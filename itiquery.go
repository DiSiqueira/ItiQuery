package itiquery

import "net/http"

// QueryMatcher Store the schema to match with the request Path
type QueryMatcher struct {
	name  string
	value string
}

// New is the constructor to ItiQuery
func New(header, value string) *QueryMatcher {
	return &QueryMatcher{
		name:  header,
		value: value,
	}
}

// Match returns if the request can be handled by this Route.
func (q *QueryMatcher) Match(req *http.Request) bool {
	if req.URL.Query().Get(q.name) != q.value {
		return false
	}
	return true
}
