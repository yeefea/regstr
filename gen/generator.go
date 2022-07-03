package gen

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"unicode"
)

const (
	charNL = '\n'
)

// StringGenerator
type StringGenerator interface {
	Gen() (string, error)
}

type ConcatGen struct {
	SubGens []StringGenerator
}

func (g *ConcatGen) AddSubGenerator(sub StringGenerator) {
	g.SubGens = append(g.SubGens, sub)
}

func (g *ConcatGen) Gen() (string, error) {
	if len(g.SubGens) == 0 {
		return "", nil
	}
	// concatenate strings
	sb := strings.Builder{}
	for _, x := range g.SubGens {
		tmp, err := x.Gen()
		if err != nil {
			return "", err
		}
		sb.WriteString(tmp)
	}
	return sb.String(), nil
}

type AnyCharNoNLGen struct {
}

func (g *AnyCharNoNLGen) Gen() (string, error) {
	for {
		ch := rand.Int31n(unicode.MaxRune)
		if ch != charNL {
			return fmt.Sprintf("%c", ch), nil
		}
	}
}

type LiteralGen struct {
	Literal string
}

func (g *LiteralGen) Gen() (string, error) {
	return g.Literal, nil
}

//CharClassGen handles [a-zA-Z0-9]
type CharClassGen struct {
	Ranges []rune
}

func (g *CharClassGen) Gen() (string, error) {
	var end int
	if len(g.Ranges)%2 == 0 {
		end = len(g.Ranges)
	} else {
		end = len(g.Ranges) - 1
	}
	if end == 0 {
		return "", errors.New("no viable rule")
	}

	var cumulation rune = 0
	offsetList := make([]rune, 0)
	idxList := make([]rune, 0)
	for idx := 0; idx < end; idx += 2 {
		offsetList = append(offsetList, g.Ranges[idx]-cumulation)
		idxList = append(idxList, cumulation)
		cumulation += g.Ranges[idx+1] - g.Ranges[idx] + 1

	}
	x := rand.Int31n(cumulation)
	offsetIdx := binarySearch(idxList, x)
	offset := offsetList[offsetIdx]
	return fmt.Sprintf("%c", x+offset), nil
}

// RepeatGen handles ? + * {m,n}
type RepeatGen struct {
	Sub StringGenerator
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

	sb := strings.Builder{}
	for i := 0; i < nTimes; i++ {
		tmp, err := g.Sub.Gen()
		if err != nil {
			return "", err
		}
		sb.WriteString(tmp)
	}
	return sb.String(), nil

}

// EmptyGen handles empty string
type EmptyGen struct {
}

func (g *EmptyGen) Gen() (string, error) {
	return "", nil
}

// AlternateGen handles |
type AlternateGen struct {
	*ConcatGen
}

func (g *AlternateGen) Gen() (string, error) {
	len := len(g.SubGens)
	if len == 0 {
		return "", errors.New("no viable rule")
	}
	// randomly select a generator
	return g.SubGens[rand.Intn(len)].Gen()
}
