package gen

import (
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

type AnyCharNoNLGen struct {
	*BaseGen
}

func (g *AnyCharNoNLGen) Gen() (string, error) {
	return string(charset[rand.Intn(lenCharset)]), nil
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

type RepeatGen struct {
	*BaseGen
	Min int
	Max int
}

type EmptyGen struct {
	*BaseGen
}

func (g *EmptyGen) Gen() (string, error) {
	return "", nil
}
