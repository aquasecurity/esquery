package esquery

import (
	"testing"
)

func TestHighlight(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"simple highlight",
			Highlight().Field("content"),
			map[string]interface{}{
				"fields": map[string]interface{}{
					"content": map[string]interface{}{},
				},
			},
		},
		{
			"highlight all params",
			Highlight().
				PreTags("<pre>", "<code>").
				PostTags("</code>", "</pre>").
				Field("content",
					Highlight().
						BoundaryChars(".;,")).
				FragmentSize(150).
				NumberOfFragments(4).
				Type(HighlighterPlain).
				BoundaryChars("()[]").
				BoundaryMaxScan(32).
				BoundaryScanner(BoundaryScannerChars).
				BoundaryScannerLocale("en-US").
				Encoder(EncoderHtml).
				ForceSource(true).
				Fragmenter(FragmenterSimple).
				FragmentOffset(6).
				HighlightQuery(
					Bool().
						Must(
							Match("author").
								Query("some guy").
								Analyzer("analyzer?").
								Fuzziness("fuzz"))).
				MatchedFields("title", "body").
				NoMatchSize(64).
				Order(OrderScore).
				PhraseLimit(512).
				RequireFieldMatch(false).
				TagsSchema(TagsSchemaStyled),
			map[string]interface{}{
				"pre_tags":                []string{"<pre>", "<code>"},
				"post_tags":               []string{"</code>", "</pre>"},
				"fragment_size":           150,
				"number_of_fragments":     4,
				"type":                    "plain",
				"boundary_chars":          "()[]",
				"boundary_scanner":        "chars",
				"boundary_max_scan":       32,
				"boundary_scanner_locale": "en-US",
				"encoder":                 "html",
				"force_source":            true,
				"fragment_offset":         6,
				"fragmenter":              "simple",
				"matched_fields":          []string{"title", "body"},
				"no_match_size":           64,
				"order":                   "score",
				"phrase_limit":            512,
				"require_field_match":     false,
				"tags_schema":             "styled",
				"fields": map[string]interface{}{
					"content": map[string]interface{}{
						"boundary_chars": ".;,",
					},
				},
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []map[string]interface{}{
							{
								"match": map[string]interface{}{
									"author": map[string]interface{}{
										"analyzer":  "analyzer?",
										"fuzziness": "fuzz",
										"query":     "some guy",
									},
								},
							},
						},
					},
				},
			},
		},
	})
}
