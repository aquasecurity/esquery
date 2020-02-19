# esquery

**esquery** is a non-obtrusive, idiomatic and easy-to-use query and aggregation builder for the [official Go client](https://github.com/elastic/go-elasticsearch) for [ElasticSearch](https://www.elastic.co/products/elasticsearch). It alleviates the need to use extremely nested maps (`map[string]interface{}`) and serializing queries to JSON manually. It also helps eliminating common mistakes such as misspelling query types, as everything is statically typed.

Save yourself some joint aches and many lines of code by switching for maps to `esquery`. Wanna know how much code you'll save? just read this project's test.

## Usage

esquery provides a [method chaining](https://en.wikipedia.org/wiki/Method_chaining)-style API for building and executing queries and aggregations. It does not wrap the official Go client nor does it require you to change your existing code in order to integrate the library. Queries can be directly built with `esquery`, and executed by passing an `*elasticsearch.Client` instance (with optional search parameters). Results are returned as-is from the official client (e.g. `*esapi.Response` objects).

Getting started is extremely simple:

```go
package main

import (
	"context"
	"log"

	"bitbucket.org/scalock/esquery"
	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
    // connect to an ElasticSearch instance
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Failed creating client: %s", err)
	}

    // run a boolean search query
	qRes, err := esquery.Query(
		esquery.
			Bool().
			Must(esquery.Term("title", "Go and Stuff")).
			Filter(esquery.Term("tag", "tech")),
    ).Run(
        es, 
		es.Search.WithContext(context.TODO()),
		es.Search.WithIndex("test"),
	)
	if err != nil {
		log.Fatalf("Failed searching for stuff: %s", err)
	}

	defer qRes.Body.Close()

	// run an aggregation
	aRes, err := esquery.Aggregate(
		esquery.Avg("average_score", "score"),
		esquery.Max("max_score", "score"),
	).Run(
		es,
		es.Search.WithContext(context.TODO()),
		es.Search.WithIndex("test"),
	)
	if err != nil {
		log.Fatalf("Failed searching for stuff: %s", err)
	}

	defer aRes.Body.Close()

    // ...
}
```

## Notes

* `esquery` currently supports version 7 of the ElasticSearch Go client.
* The library cannot currently generate "short queries". For example, whereas
  ElasticSearch can accept this:

```json
{ "query": { "term": { "user": "Kimchy" } } }
```

  The library will always generate this:

```json
{ "query": { "term": { "user": { "value": "Kimchy" } } } }
```

  This is also true for queries such as "bool", where fields like "must" can
  either receive one query object, or an array of query objects. `esquery` will
  generate an array even if there's only one query object.

## Supported Queries

The following queries are currently supported:

| ElasticSearch DSL       | `esquery` Function    |
| ------------------------|---------------------- |
| `"match"`               | `Match()`             |
| `"match_bool_prefix"`   | `MatchBoolPrefix()`   |
| `"match_phrase"`        | `MatchPhrase()`       |
| `"match_phrase_prefix"` | `MatchPhrasePrefix()` |
| `"match_all"`           | `MatchAll()`          |
| `"match_none"`          | `MatchNone()`         |
| `"exists"`              | `Exists()`            |
| `"fuzzy"`               | `Fuzzy()`             |
| `"ids"`                 | `IDs()`               |
| `"prefix"`              | `Prefix()`            |
| `"range"`               | `Range()`             |
| `"regexp"`              | `Regexp()`            |
| `"term"`                | `Term()`              |
| `"terms"`               | `Terms()`             |
| `"terms_set"`           | `TermsSet()`          |
| `"wildcard"`            | `Wildcard()`          |
| `"bool"`                | `Bool()`              |
| `"boosting"`            | `Boosting()`          |
| `"constant_score"`      | `ConstantScore()`     |
| `"dis_max"`             | `DisMax()`            |

### Custom Queries

To execute an arbitrary query, or any query that is not natively supported by the library yet, use the `CustomQuery()` function, which accepts any `map[string]interface{}` value.

## Supported Aggregations

The following aggregations are currently supported:

| ElasticSearch DSL       | `esquery` Function    |
| ------------------------|---------------------- |
| `"avg"`                 | `Avg()`               |
| `"weighted_avg"`        | `WeightedAvg()`       |
| `"cardinality"`         | `Cardinality()`       |
| `"max"`                 | `Max()`               |
| `"min"`                 | `Min()`               |
| `"sum"`                 | `Sum()`               |
| `"value_count"`         | `ValueCount()`        |
| `"percentiles"`         | `Percentiles()`       |
| `"stats"`               | `Stats()`             |
| `"string_stats"`        | `StringStats()`       |

### Custom Aggregations

To execute an arbitrary aggregation, or any aggregation that is not natively supported by the library yet, use the `CustomAgg()` function, which accepts any `map[string]interface{}` value.
