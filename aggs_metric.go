package esquery

import "github.com/fatih/structs"

type BaseAgg struct {
	name           string
	apiName        string
	*BaseAggParams `structs:",flatten"`
}

type BaseAggParams struct {
	Field string      `structs:"field"`
	Miss  interface{} `structs:"missing,omitempty"`
}

func newBaseAgg(apiName, name, field string) *BaseAgg {
	return &BaseAgg{
		name:    name,
		apiName: apiName,
		BaseAggParams: &BaseAggParams{
			Field: field,
		},
	}
}

func (agg *BaseAgg) Name() string {
	return agg.name
}

func (agg *BaseAgg) Map() map[string]interface{} {
	return map[string]interface{}{
		agg.apiName: structs.Map(agg.BaseAggParams),
	}
}

/*******************************************************************************
 * Avg Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-avg-aggregation.html
 ******************************************************************************/

type AvgAgg struct {
	*BaseAgg `structs:",flatten"`
}

func Avg(name, field string) *AvgAgg {
	return &AvgAgg{
		BaseAgg: newBaseAgg("avg", name, field),
	}
}

func (agg *AvgAgg) Missing(val interface{}) *AvgAgg {
	agg.Miss = val
	return agg
}

/*******************************************************************************
 * Weighed Avg Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-weight-avg-aggregation.html
 ******************************************************************************/

type WeightedAvgAgg struct {
	name    string
	apiName string
	Val     *BaseAggParams `structs:"value"`
	Weig    *BaseAggParams `structs:"weight"`
}

func WeightedAvg(name string) *WeightedAvgAgg {
	return &WeightedAvgAgg{
		name:    name,
		apiName: "weighted_avg",
	}
}

func (agg *WeightedAvgAgg) Name() string {
	return agg.name
}

func (agg *WeightedAvgAgg) Value(field string, missing ...interface{}) *WeightedAvgAgg {
	agg.Val = new(BaseAggParams)
	agg.Val.Field = field
	if len(missing) > 0 {
		agg.Val.Miss = missing[len(missing)-1]
	}
	return agg
}

func (agg *WeightedAvgAgg) Weight(field string, missing ...interface{}) *WeightedAvgAgg {
	agg.Weig = new(BaseAggParams)
	agg.Weig.Field = field
	if len(missing) > 0 {
		agg.Weig.Miss = missing[len(missing)-1]
	}
	return agg
}

func (agg *WeightedAvgAgg) Map() map[string]interface{} {
	return map[string]interface{}{
		agg.apiName: structs.Map(agg),
	}
}

/*******************************************************************************
 * Cardinality Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-cardinality-aggregation.html
 ******************************************************************************/

type CardinalityAgg struct {
	*BaseAgg     `structs:",flatten"`
	PrecisionThr uint16 `structs:"precision_threshold,omitempty"`
}

func Cardinality(name, field string) *CardinalityAgg {
	return &CardinalityAgg{
		BaseAgg: newBaseAgg("cardinality", name, field),
	}
}

func (agg *CardinalityAgg) Missing(val interface{}) *CardinalityAgg {
	agg.Miss = val
	return agg
}

func (agg *CardinalityAgg) PrecisionThreshold(val uint16) *CardinalityAgg {
	agg.PrecisionThr = val
	return agg
}

func (agg *CardinalityAgg) Map() map[string]interface{} {
	return map[string]interface{}{
		agg.apiName: structs.Map(agg),
	}
}

/*******************************************************************************
 * Max Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-max-aggregation.html
 ******************************************************************************/

type MaxAgg struct {
	*BaseAgg `structs:",flatten"`
}

func Max(name, field string) *MaxAgg {
	return &MaxAgg{
		BaseAgg: newBaseAgg("max", name, field),
	}
}

func (agg *MaxAgg) Missing(val interface{}) *MaxAgg {
	agg.Miss = val
	return agg
}

/*******************************************************************************
 * Min Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-min-aggregation.html
 ******************************************************************************/

type MinAgg struct {
	*BaseAgg `structs:",flatten"`
}

func Min(name, field string) *MinAgg {
	return &MinAgg{
		BaseAgg: newBaseAgg("min", name, field),
	}
}

func (agg *MinAgg) Missing(val interface{}) *MinAgg {
	agg.Miss = val
	return agg
}

/*******************************************************************************
 * Sum Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-sum-aggregation.html
 ******************************************************************************/

type SumAgg struct {
	*BaseAgg `structs:",flatten"`
}

func Sum(name, field string) *SumAgg {
	return &SumAgg{
		BaseAgg: newBaseAgg("sum", name, field),
	}
}

func (agg *SumAgg) Missing(val interface{}) *SumAgg {
	agg.Miss = val
	return agg
}

/*******************************************************************************
 * Value Count Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-valuecount-aggregation.html
 ******************************************************************************/

type ValueCountAgg struct {
	*BaseAgg `structs:",flatten"`
}

func ValueCount(name, field string) *ValueCountAgg {
	return &ValueCountAgg{
		BaseAgg: newBaseAgg("value_count", name, field),
	}
}

/*******************************************************************************
 * Percentiles Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-percentile-aggregation.html
 ******************************************************************************/

type PercentilesAgg struct {
	*BaseAgg `structs:",flatten"`
	Prcnts   []float32 `structs:"percents,omitempty"`
	Key      *bool     `structs:"keyed,omitempty"`
	TDigest  struct {
		Compression uint16 `structs:"compression,omitempty"`
	} `structs:"tdigest,omitempty"`
	HDR struct {
		NumHistogramDigits uint8 `structs:"number_of_significant_value_digits,omitempty"`
	} `structs:"hdr,omitempty"`
}

func Percentiles(name, field string) *PercentilesAgg {
	return &PercentilesAgg{
		BaseAgg: newBaseAgg("percentiles", name, field),
	}
}

func (agg *PercentilesAgg) Percents(percents ...float32) *PercentilesAgg {
	agg.Prcnts = percents
	return agg
}

func (agg *PercentilesAgg) Missing(val interface{}) *PercentilesAgg {
	agg.Miss = val
	return agg
}

func (agg *PercentilesAgg) Keyed(b bool) *PercentilesAgg {
	agg.Key = &b
	return agg
}

func (agg *PercentilesAgg) Compression(val uint16) *PercentilesAgg {
	agg.TDigest.Compression = val
	return agg
}

func (agg *PercentilesAgg) NumHistogramDigits(val uint8) *PercentilesAgg {
	agg.HDR.NumHistogramDigits = val
	return agg
}

func (agg *PercentilesAgg) Map() map[string]interface{} {
	return map[string]interface{}{
		agg.apiName: structs.Map(agg),
	}
}

/*******************************************************************************
 * Stats Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-stats-aggregation.html
 ******************************************************************************/

type StatsAgg struct {
	*BaseAgg `structs:",flatten"`
}

func Stats(name, field string) *StatsAgg {
	return &StatsAgg{
		BaseAgg: newBaseAgg("stats", name, field),
	}
}

func (agg *StatsAgg) Missing(val interface{}) *StatsAgg {
	agg.Miss = val
	return agg
}

/*******************************************************************************
 * String Stats Aggregation
 * https://www.elastic.co/guide/en/elasticsearch/reference/
 *    current/search-aggregations-metrics-string-stats-aggregation.html
 ******************************************************************************/

type StringStatsAgg struct {
	*BaseAgg `structs:",flatten"`
	ShowDist *bool `structs:"show_distribution,omitempty"`
}

func StringStats(name, field string) *StringStatsAgg {
	return &StringStatsAgg{
		BaseAgg: newBaseAgg("string_stats", name, field),
	}
}

func (agg *StringStatsAgg) Missing(val interface{}) *StringStatsAgg {
	agg.Miss = val
	return agg
}

func (agg *StringStatsAgg) ShowDistribution(b bool) *StringStatsAgg {
	agg.ShowDist = &b
	return agg
}

func (agg *StringStatsAgg) Map() map[string]interface{} {
	return map[string]interface{}{
		agg.apiName: structs.Map(agg),
	}
}
