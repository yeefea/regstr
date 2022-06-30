package regstr

import (
	"errors"
	"fmt"
	"regexp/syntax"

	"github.com/yeefea/regstr/gen"
)

func MustCompile(s string) gen.StringGenerator {
	r, err := syntax.Parse(s, syntax.Perl)
	if err != nil {
		panic(err)
	}
	g, err := newStringGeneratorByRegexp(r)
	if err != nil {
		panic(err)
	}
	return g
}

func Compile(s string) (gen.StringGenerator, error) {
	r, err := syntax.Parse(s, syntax.Perl)
	if err != nil {
		return nil, err
	}
	return newStringGeneratorByRegexp(r)
}

func newStringGeneratorByRegexp(x *syntax.Regexp) (gen.StringGenerator, error) {
	// The following empty strings are not supported.
	// syntax.OpBeginLine matches empty string at beginning of line
	// syntax.OpEndLine matches empty string at end of line
	// syntax.OpBeginText matches empty string at beginning of text
	// syntax.OpEndText matches empty string at end of text
	// syntax.OpWordBoundary matches word boundary `\b`
	// syntax.OpNoWordBoundary matches word non-boundary `\B`

	// i s m U flags are not supported.
	// syntax.OpAnyChar matches any character

	var g gen.StringGenerator
	var err error
	switch x.Op {
	case syntax.OpNoMatch:
		// matches no strings
		return nil, errors.New("matches no strings")
	case syntax.OpEmptyMatch:
		// matches empty string
		g = &gen.EmptyGen{}
	case syntax.OpLiteral:
		// matches Runes sequence
		g = &gen.LiteralGen{BaseGen: &gen.BaseGen{}, Text: string(x.Rune)}
	case syntax.OpCharClass:
		fmt.Println("TODO") // matches Runes interpreted as range pair list
	case syntax.OpAnyCharNotNL:
		// matches any character except newline
		g = &gen.AnyCharNoNLGen{BaseGen: &gen.BaseGen{}}
	case syntax.OpCapture:
		// capturing subexpression with index Cap, optional name Name
		tmp := &gen.BaseGen{}
		for _, sub := range x.Sub {
			subG, err := newStringGeneratorByRegexp(sub)
			if err != nil {
				return nil, err
			}
			tmp.AddSubGenerator(subG)
		}
		g = tmp
	case syntax.OpStar:
		fmt.Println("TODO") // matches Sub[0] zero or more times
	case syntax.OpPlus:
		// matches Sub[0] one or more times
		g, err = newRepeatGen(x.Sub[0], 1, 10)
		if err != nil {
			return nil, err
		}
	case syntax.OpQuest:
		// matches Sub[0] zero or one times
		tmp := &gen.RepeatGen{BaseGen: &gen.BaseGen{}, Min: 0, Max: 1}
		subG, err := newStringGeneratorByRegexp(x.Sub[0])
		if err != nil {
			return nil, err
		}
		tmp.AddSubGenerator(subG)
		g = tmp
	case syntax.OpRepeat:
		// matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
		var err error
		g, err = newRepeatGen(x.Sub[0], x.Min, x.Max)
		if err != nil {
			return nil, err
		}
	case syntax.OpConcat:
		// matches concatenation of Subs
		tmp := &gen.BaseGen{}
		for _, sub := range x.Sub {
			subG, err := newStringGeneratorByRegexp(sub)
			if err != nil {
				return nil, err
			}
			tmp.AddSubGenerator(subG)
		}
		g = tmp
	case syntax.OpAlternate:
		// matches alternation of Subs
		tmp := &gen.AlternateGen{BaseGen: &gen.BaseGen{}}
		for _, sub := range x.Sub {
			subG, err := newStringGeneratorByRegexp(sub)
			if err != nil {
				return nil, err
			}
			tmp.AddSubGenerator(subG)
		}
		g = tmp
	}
	return g, nil
}

func newRepeatGen(sub0 *syntax.Regexp, min int, max int) (*gen.RepeatGen, error) {
	g := &gen.RepeatGen{BaseGen: &gen.BaseGen{}, Min: min, Max: max}
	subG, err := newStringGeneratorByRegexp(sub0)
	if err != nil {
		return nil, err
	}
	g.AddSubGenerator(subG)
	return g, nil
}
