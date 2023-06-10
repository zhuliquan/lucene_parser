package lucene_parser

import (
	"github.com/zhuliquan/lucene_parser/term"
)

func TermGroupToLucene(field *term.Field, termGroup *term.TermGroup) *Lucene {
	if termGroup == nil {
		return nil
	} else {
		return logicTermGroupToLucene(field, termGroup.LogicTermGroup, termGroup.BoostSymbol)
	}
}

func logicTermGroupToLucene(field *term.Field, termGroup *term.LogicTermGroup, boostSymbol string) *Lucene {
	if termGroup == nil {
		return nil
	} else {
		var q = &Lucene{}
		q.OrQuery = orTermGroupToOrQuery(field, termGroup.OrTermGroup, boostSymbol)
		if q.OrQuery == nil {
			return nil
		}
		for _, osGroup := range termGroup.OSTermGroup {
			if t := osTermGroupToOsQuery(field, osGroup, boostSymbol); t != nil {
				q.OSQuery = append(q.OSQuery, t)
			}
		}
		return q
	}
}

func osTermGroupToOsQuery(field *term.Field, osGroup *term.OSTermGroup, boostSymbol string) *OSQuery {
	if osGroup == nil {
		return nil
	} else {
		if t := orTermGroupToOrQuery(field, osGroup.OrTermGroup, boostSymbol); t != nil {
			return &OSQuery{
				OrSymbol: osGroup.OrSymbol,
				OrQuery:  t,
			}
		} else {
			return nil
		}
	}
}

func orTermGroupToOrQuery(field *term.Field, group *term.OrTermGroup, boostSymbol string) *OrQuery {
	if group == nil {
		return nil
	} else {
		var q = &OrQuery{}
		q.AndQuery = andGroupToAndQuery(field, group.AndTermGroup, boostSymbol)
		if q.AndQuery == nil {
			return nil
		}

		for _, ansGroup := range group.AnSTermGroup {
			if t := ansGroupToAnsQuery(field, ansGroup, boostSymbol); t != nil {
				q.AnSQuery = append(q.AnSQuery, t)
			}
		}
		return q
	}
}

func ansGroupToAnsQuery(field *term.Field, group *term.AnSTermGroup, boostSymbol string) *AnSQuery {
	if group == nil {
		return nil
	} else {
		if t := andGroupToAndQuery(field, group.AndTermGroup, boostSymbol); t != nil {
			return &AnSQuery{
				AndSymbol: group.AndSymbol,
				AndQuery:  t,
			}
		} else {
			return nil
		}
	}
}

func andGroupToAndQuery(field *term.Field, group *term.AndTermGroup, boostSymbol string) *AndQuery {
	if group == nil {
		return nil
	}

	var fieldQuery *FieldQuery
	var parenQuery *ParenQuery

	if group.FieldTermGroup != nil {
		fieldQuery = fieldTermGroupToFieldTerm(field, group.FieldTermGroup, boostSymbol)
	} else if group.ParenTermGroup != nil {
		parenQuery = parenTermGroupToParenTerm(field, group.ParenTermGroup, boostSymbol)
	} else {
		return nil
	}
	if fieldQuery != nil || parenQuery != nil {
		return &AndQuery{
			NotSymbol:  group.NotSymbol,
			FieldQuery: fieldQuery,
			ParenQuery: parenQuery,
		}
	} else {
		return nil
	}
}

func parenTermGroupToParenTerm(field *term.Field, group *term.ParenTermGroup, boostSymbol string) *ParenQuery {
	t := logicTermGroupToLucene(field, group.SubTermGroup, boostSymbol)
	if t == nil {
		return nil
	}
	return &ParenQuery{SubQuery: t}
}

func fieldTermGroupToFieldTerm(field *term.Field, group *term.FieldTermGroup, boostSymbol string) *FieldQuery {
	// it's impossible for groupElem is nil
	if group.SingleTerm != nil {
		return &FieldQuery{
			Field: field,
			Term: &term.Term{
				FuzzyTerm: &term.FuzzyTerm{
					SingleTerm:  group.SingleTerm,
					BoostSymbol: boostSymbol,
				},
			},
		}
	} else if group.PhraseTerm != nil {
		return &FieldQuery{
			Field: field,
			Term: &term.Term{
				FuzzyTerm: &term.FuzzyTerm{
					PhraseTerm:  group.PhraseTerm,
					BoostSymbol: boostSymbol,
				},
			},
		}
	} else if group.SRangeTerm != nil {
		return &FieldQuery{
			Field: field,
			Term: &term.Term{
				RangeTerm: &term.RangeTerm{
					SRangeTerm:  group.SRangeTerm,
					BoostSymbol: boostSymbol,
				},
			},
		}
	} else if group.DRangeTerm != nil {
		return &FieldQuery{
			Field: field,
			Term: &term.Term{
				RangeTerm: &term.RangeTerm{
					DRangeTerm:  group.DRangeTerm,
					BoostSymbol: boostSymbol,
				},
			},
		}
	} else {
		return nil
	}
}
