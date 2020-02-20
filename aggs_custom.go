package esquery

type CustomAggregation struct {
	m map[string]interface{}
}

func CustomAgg(m map[string]interface{}) *CustomAggregation {
	return &CustomAggregation{m}
}

func (agg *CustomAggregation) Map() map[string]interface{} {
	return agg.m
}
