package gen

import (
	"errors"
	"regexp/syntax"
)

// StringGeneratorFactory
type StringGeneratorFactory struct {
	Limit int
}

func (f *StringGeneratorFactory) NewStringGenerator(x *syntax.Regexp) (StringGenerator, error) {
	// The following empty strings are not supported.
	// syntax.OpBeginLine matches empty string at beginning of line
	// syntax.OpEndLine matches empty string at end of line
	// syntax.OpBeginText matches empty string at beginning of text
	// syntax.OpEndText matches empty string at end of text
	// syntax.OpWordBoundary matches word boundary `\b`
	// syntax.OpNoWordBoundary matches word non-boundary `\B`

	// i s m U flags are not supported.
	// syntax.OpAnyChar matches any character

	var g StringGenerator
	var err error
	switch x.Op {
	case syntax.OpNoMatch:
		// matches no strings
		return nil, errors.New("matches no strings")
	case syntax.OpEmptyMatch:
		// matches empty string
		g = &EmptyGen{}
	case syntax.OpLiteral:
		// matches Runes sequence
		g = &LiteralGen{Literal: string(x.Rune)}
	case syntax.OpCharClass:
		// matches Runes interpreted as range pair list
		g = &CharClassGen{Ranges: x.Rune}
	case syntax.OpAnyCharNotNL:
		// matches any character except newline
		g = &AnyCharNoNLGen{}
	case syntax.OpCapture:
		// capturing subexpression with index Cap, optional name Name
		tmp := &ConcatGen{}
		for _, sub := range x.Sub {
			subG, err := f.NewStringGenerator(sub)
			if err != nil {
				return nil, err
			}
			tmp.AddSubGenerator(subG)
		}
		g = tmp
	case syntax.OpStar:
		// matches Sub[0] zero or more times
		g, err = f.newRepeatGen(x.Sub[0], 0, f.Limit)
		if err != nil {
			return nil, err
		}
	case syntax.OpPlus:
		// matches Sub[0] one or more times
		g, err = f.newRepeatGen(x.Sub[0], 1, 1+f.Limit)
		if err != nil {
			return nil, err
		}
	case syntax.OpQuest:
		// matches Sub[0] zero or one times
		subG, err := f.NewStringGenerator(x.Sub[0])
		if err != nil {
			return nil, err
		}
		tmp := &RepeatGen{Sub: subG, Min: 0, Max: 1}
		g = tmp
	case syntax.OpRepeat:
		// matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
		var err error

		g, err = f.newRepeatGen(x.Sub[0], x.Min, x.Max)
		if err != nil {
			return nil, err
		}
	case syntax.OpConcat:
		// matches concatenation of Subs
		tmp := &ConcatGen{}
		for _, sub := range x.Sub {
			subG, err := f.NewStringGenerator(sub)
			if err != nil {
				return nil, err
			}
			tmp.AddSubGenerator(subG)
		}
		g = tmp
	case syntax.OpAlternate:
		// matches alternation of Subs
		tmp := &AlternateGen{ConcatGen: &ConcatGen{}}
		for _, sub := range x.Sub {
			subG, err := f.NewStringGenerator(sub)
			if err != nil {
				return nil, err
			}
			tmp.AddSubGenerator(subG)
		}
		g = tmp
	default:
		g = &EmptyGen{}
	}
	return g, nil
}

func (f *StringGeneratorFactory) newRepeatGen(sub0 *syntax.Regexp, min int, max int) (*RepeatGen, error) {

	subG, err := f.NewStringGenerator(sub0)
	if err != nil {
		return nil, err
	}
	g := &RepeatGen{Sub: subG, Min: min, Max: max}
	return g, nil
}
