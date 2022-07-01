package main

import (
	"fmt"
	"testing"

	"github.com/yeefea/regstr/gen"
)

func TestRepeatGen(t *testing.T) {
	base := gen.ConcatGen{}
	base.AddSubGenerator(&gen.LiteralGen{Literal: "123"})
	g := gen.RepeatGen{Sub: &base, Min: 1, Max: 1}
	fmt.Println(g.Gen())
}
