package standard

import (
	"encoding/json"
	"testing"
)

func TestYesYacc(t *testing.T) {
	ShowGrammar()
	type cases struct {
		name string
		args string
	}

	for _, tt := range []cases{
		{
			name: "boolean-1",
			args: `"jakarta apache" jakarta`,
		},
		{
			name: "boolean-2",
			args: `"jakarta apache" OR jakarta`,
		},
		{
			name: "boolean-3",
			args: `"jakarta apache" or jakarta`,
		},
		{
			name: "boolean-4",
			args: `"jakarta apache" AND "Apache Lucene"`,
		},
		{
			name: "boolean-5",
			args: `+jakarta lucene`,
		},
		{
			name: "boolean-6",
			args: `"jakarta apache" NOT "Apache Lucene"`,
		},
		{
			name: "boolean-7",
			args: `NOT "jakarta apache"`,
		},
		{
			name: "boolean-8",
			args: `!"jakarta apache"`,
		},
		{
			name: "boolean-9",
			args: `"jakarta apache" -"Apache Lucene"`,
		},
		{
			name: "boost-1",
			args: `jakarta apache`,
		},
		{
			name: "boost-2",
			args: `jack^ micheal`,
		},
		{
			name: "boost-3",
			args: `"jakarta apache"^4 "Apache Lucene"`,
		},
		{
			name: "escape-1",
			args: `\(1\+1\)\:2`,
		},
		{
			name: "side-range-1",
			args: `(+>1 +<10)^10`,
		},
		{
			name: "double-range-1",
			args: `[ "das" TO * }`,
		},
		{
			name: "double_range-2",
			args: `title:{Aida TO Carmen}`,
		},
		{
			name: "field-1",
			args: `title:"The Right Way" AND text:go`,
		},
		{
			name: "field-2",
			args: `title:"Do it right" AND right`,
		},
		{
			name: "field-3",
			args: `title:Do it right`,
		},
		{
			name: "fuzzy-1",
			args: `roam~`,
		},
		{
			name: "fuzzy-2",
			args: `roam~0.8`,
		},
		{
			name: "fuzzy-3",
			args: `"jakarta apache"~10`,
		},
		{
			name: "group-1",
			args: `(jakarta OR apache) AND website`,
		},
		{
			name: "group-2",
			args: `title:(+return +"pink panther")`,
		},
		{
			name: "wildcard-1",
			args: `te?t`,
		},
		{
			name: "wildcard-2",
			args: `test*`,
		},
		{
			name: "wildcard-3",
			args: `tes*t`,
		},
		{
			name: "regex_01",
			args: `/[0-9]+(\.[0-9])?/`,
		},
		{
			name: "regex_02",
			args: `age:/[0-9]+(\.[0-9])?/`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			qry, err := ParseLucene(tt.args)
			if err != nil {
				t.Logf("expect not error, but got: %+v", err)
				t.Fail()
			} else {
				b, _ := json.Marshal(qry)
				t.Logf("success to got: %s", b)
			}
		})
	}

}
