package gen

import (
	"errors"
	"math/rand"
	"strings"
)

const (
	maxCount   = 16
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	lenCharset = len(charset)
)

type StringGenerator interface {
	Gen() (string, error)
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

type AnyCharNoNLGen struct {
	*BaseGen
}

func (g *AnyCharNoNLGen) Gen() (string, error) {
	return string(charset[rand.Intn(lenCharset)]), nil
}

type LiteralGen struct {
	*BaseGen
	Text string
}

func (g *LiteralGen) Gen() (string, error) {
	return g.Text, nil
}

type CharClassGen struct {
	*BaseGen
	Ranges []int
}

type RepeatGen struct {
	*BaseGen
	Min int
	Max int
}

func (g *RepeatGen) Gen() (string, error) {
	var nTimes int
	if g.Min < g.Max {
		nTimes = rand.Intn(g.Max-g.Min+1) + g.Min
	} else if g.Min == 0 {
		return "", nil
	} else {
		nTimes = g.Min
	}
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

type EmptyGen struct {
}

func (g *EmptyGen) Gen() (string, error) {
	return "", nil
}

type AlternateGen struct {
	*BaseGen
}

func (g *AlternateGen) Gen() (string, error) {
	len := len(g.BaseGen.SubGens)
	if len == 0 {
		return "", errors.New("no viable rule")
	}
	return g.SubGens[rand.Intn(len)].Gen()
}
