package regstr

import (
	"errors"
	"fmt"
	"regexp/syntax"

	"github.com/yeefea/regstr/gen"
)

func Parse(s string, flags syntax.Flags) (gen.StringGenerator, error) {
	r, err := syntax.Parse(s, flags)
	if err != nil {
		return nil, err
	}
	return traverseRegexp(r)
}

func traverseRegexp(x *syntax.Regexp) (gen.StringGenerator, error) {
	var g gen.StringGenerator
	switch x.Op {
	case syntax.OpNoMatch:
		// matches no strings
		return nil, errors.New("matches no strings")
	case syntax.OpEmptyMatch:
		// matches empty string
		g = &gen.EmptyGen{BaseGen: &gen.BaseGen{}}
	case syntax.OpLiteral:
		// matches Runes sequence
		g = &gen.FixedGen{BaseGen: &gen.BaseGen{}, Text: string(x.Rune)}
	case syntax.OpCharClass:
		fmt.Println("TODO") // matches Runes interpreted as range pair list
	case syntax.OpAnyCharNotNL:
		// matches any character except newline
		g = &gen.AnyCharNoNLGen{BaseGen: &gen.BaseGen{}}
	case syntax.OpAnyChar:
		fmt.Println("TODO") // matches any character
	case syntax.OpBeginLine:
		fmt.Println("TODO") // matches empty string at beginning of line
	case syntax.OpEndLine:
		fmt.Println("TODO") // matches empty string at end of line
	case syntax.OpBeginText:
		fmt.Println("TODO") // matches empty string at beginning of text
	case syntax.OpEndText:
		fmt.Println("TODO") // matches empty string at end of text
	case syntax.OpWordBoundary:
		fmt.Println("TODO") // matches word boundary `\b`
	case syntax.OpNoWordBoundary:
		fmt.Println("TODO") // matches word non-boundary `\B`
	case syntax.OpCapture:
		fmt.Println("TODO") // capturing subexpression with index Cap, optional name Name
	case syntax.OpStar:
		fmt.Println("TODO") // matches Sub[0] zero or more times
	case syntax.OpPlus:
		// matches Sub[0] one or more times
		g = &gen.PlusGen{BaseGen: &gen.BaseGen{}}
		subG, err := traverseRegexp(x.Sub[0])
		if err != nil {
			return nil, err
		}
		g.AddSubGenerator(subG)
	case syntax.OpQuest:
		// matches Sub[0] zero or one times
		g = &gen.RepeatGen{BaseGen: &gen.BaseGen{}, Min: 0, Max: 1}
		subG, err := traverseRegexp(x.Sub[0])
		if err != nil {
			return nil, err
		}
		g.AddSubGenerator(subG)
	case syntax.OpRepeat:
		// matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
		g = &gen.RepeatGen{BaseGen: &gen.BaseGen{}, Min: x.Min, Max: x.Max}
	case syntax.OpConcat:
		// matches concatenation of Subs
		g = &gen.BaseGen{}
		for _, sub := range x.Sub {
			subG, err := traverseRegexp(sub)
			if err != nil {
				return nil, err
			}
			g.AddSubGenerator(subG)
		}
	case syntax.OpAlternate:
		fmt.Println("TODO") // matches alternation of Subs
	}
	return g, nil
}
