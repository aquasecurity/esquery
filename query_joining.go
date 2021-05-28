package esquery

// NestedQuery represents a query of type nested as described in:
// https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-nested-query.html
type NestedQuery struct {
	path           string
	query          Mappable
	scoreMode      string
	ignoreUnmapped bool
}

func Nested(path string, query Mappable) *NestedQuery {
	return &NestedQuery{
		path:  path,
		query: query,
	}
}

func (n *NestedQuery) ScoreMode(mode string) *NestedQuery {
	n.scoreMode = mode
	return n
}

func (n *NestedQuery) IgnoreUnmapped(val bool) *NestedQuery {
	n.ignoreUnmapped = val
	return n
}

// Map returns a map representation of the query, thus implementing the
// Mappable interface.
func (n *NestedQuery) Map() map[string]interface{} {
	innerMap := map[string]interface{}{"path": n.path, "query": n.query.Map()}
	if n.scoreMode != "" {
		innerMap["score_mode"] = n.scoreMode
	}
	if n.ignoreUnmapped == true {
		innerMap["ignore_unmapped"] = n.ignoreUnmapped
	}
	return map[string]interface{}{"nested": innerMap}
}
