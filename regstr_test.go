package regstr

import (
	"fmt"
	"testing"

	"github.com/yeefea/regstr/gen"
)

func TestRepeatGen(t *testing.T) {
	base := gen.BaseGen{}
	base.AddSubGenerator(&gen.LiteralGen{Text: "123"})
	g := gen.RepeatGen{BaseGen: &base, Min: 1, Max: 1}
	fmt.Println(g.Gen())
}
