package esquery

type CustomQry struct {
	m map[string]interface{}
}

func CustomQuery(m map[string]interface{}) *CustomQry {
	return &CustomQry{m}
}

func (q *CustomQry) Map() map[string]interface{} {
	return q.m
}
