# esquery

`esquery` is an idiomatic, easy-to-use query builder for the [official Go client](https://github.com/elastic/go-elasticsearch) for [ElasticSearch](https://www.elastic.co/products/elasticsearch). It alleviates the need to use extremely nested maps of empty interfaces and serializing queries to JSON manually. It also helps eliminating common mistakes such as misspelling query types, as everything is statically typed.

## Usage

`esquery` can be used directly to build queries, with no need for external dependencies. It can execute the queries against an existing instance of `*esapi.API`, but the queries can also be manually converted to JSON if necessary.

```go
package main

import (
	"context"
	"log"

	"bitbucket.org/scalock/esquery"
	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Failed creating client: %s", err)
	}

	res, err := esquery.Search(
		es,
		esquery.
			Bool().
			Must(esquery.Term("title", "Go and Stuff")).
			Filter(esquery.Term("tag", "tech")),
		es.Search.WithContext(context.TODO()),
		es.Search.WithIndex("test"),
	)
	if err != nil {
		log.Fatalf("Failed searching for stuff: %s", err)
	}

	defer res.Body.Close()

	// ...
}
```

## Notes

* Library currently supports v7 of the ElasticSearch Go client.
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

## Supported queries

The following queries are currently supported:

| Query                   | `esquery` Function    |
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
