package regstr

import (
	"math/rand"
	"regexp/syntax"
	"strings"
)

const (
	maxCount   = 16
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lenCharset = len(charset)
)

func Parse(s string, flags syntax.Flags) (StringGenerator, error) {
	r, err := syntax.Parse(s, flags)
	if err != nil {
		return nil, err
	}
	g := traverseRegexp(r)
	return g, nil
}

type StringGenerator interface {
	Gen() (string, error)
	AddSubGenerator(StringGenerator)
}

type BaseGen struct {
	SubGens []StringGenerator
}

func (g *BaseGen) AddSubGenerator(sub StringGenerator) {
	g.SubGens = append(g.SubGens, sub)
}

func (g *BaseGen) Gen() (string, error) {
	if len(g.SubGens) == 0 {
		return "", nil
	}
	buf := make([]string, 0, len(g.SubGens))
	for _, x := range g.SubGens {
		tmp, err := x.Gen()
		if err != nil {
			return "", err
		}
		buf = append(buf, tmp)
	}
	return strings.Join(buf, ""), nil
}

type ConcatGen struct {
	*BaseGen
}

type AnyCharNoNLGen struct {
	*BaseGen
}

func (g *AnyCharNoNLGen) Gen() (string, error) {
	return string(charset[rand.Intn(lenCharset)]), nil
}

func (g *ConcatGen) Gen() (string, error) {
	if len(g.SubGens) == 0 {
		return "", nil
	}
	buf := make([]string, 0, len(g.SubGens))
	for _, x := range g.SubGens {
		tmp, err := x.Gen()
		if err != nil {
			return "", err
		}
		buf = append(buf, tmp)
	}
	return strings.Join(buf, ""), nil
}

type FixedGen struct {
	*BaseGen
	Text string
}

func (g *FixedGen) Gen() (string, error) {
	return g.Text, nil
}

type PlusGen struct {
	*BaseGen
}

func (g *PlusGen) Gen() (string, error) {
	nTimes := rand.Intn(maxCount-1) + 1
	buf := make([]string, 0, nTimes)
	for i := 0; i < nTimes; i++ {
		tmp, err := g.SubGens[0].Gen()
		if err != nil {
			return "", err
		}
		buf = append(buf, tmp)
	}

	return strings.Join(buf, ""), nil

}

type CharClassGen struct {
	*BaseGen
	Ranges []int
}

func traverseRegexp(x *syntax.Regexp) StringGenerator {
	var g StringGenerator
	switch x.Op {
	case syntax.OpLiteral:
		g = &FixedGen{&BaseGen{}, string(x.Rune)}
	case syntax.OpConcat:
		g = &BaseGen{}
		for _, sub := range x.Sub {
			subG := traverseRegexp(sub)
			g.AddSubGenerator(subG)
		}
	case syntax.OpAnyCharNotNL:
		g = &AnyCharNoNLGen{&BaseGen{}}
	case syntax.OpPlus:
		g = &PlusGen{&BaseGen{}}
		for _, sub := range x.Sub {
			subG := traverseRegexp(sub)
			g.AddSubGenerator(subG)
		}
	}
	// fmt.Println(x, x.Op.String(), x.Cap, x.Flags, x.Max, x.Min, x.Name, x.Rune, x.Rune0, x.Sub0)
	// fmt.Println(g)
	return g
}
