package esquery

// Source represents the "_source" option which is commonly accepted in ES
// queries. Currently, only the "includes" option is supported.
type Source struct {
	includes []string
}

// Map returns a map representation of the Source object.
func (source Source) Map() map[string]interface{} {
	return map[string]interface{}{
		"includes": source.includes,
	}
}

// Sort represents a list of keys to sort by.
type Sort []map[string]interface{}

// Order is the ordering for a sort key (ascending, descending).
type Order string

const (
	// OrderAsc represents sorting in ascending order.
	OrderAsc Order = "asc"

	// OrderDesc represents sorting in descending order.
	OrderDesc Order = "desc"
)
