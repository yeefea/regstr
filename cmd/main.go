package main

import (
	"fmt"
	"math/rand"
	"regexp/syntax"
	"time"

	"github.com/yeefea/regstr"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func DescRegexp(x *syntax.Regexp, prefix string) {
	fmt.Println(prefix, x, x.Op.String(), x.Cap, x.Flags, x.Min, x.Max, x.Name, x.Rune, x.Rune0, x.Sub0)
	for _, sub := range x.Sub {
		DescRegexp(sub, prefix+"  ")
	}
}
func demo() {
	re := `a{0,0}`
	g, err := syntax.Parse(re, syntax.Perl)
	if err != nil {
		panic(err)
	}
	DescRegexp(g, "")
	g.Simplify()
	DescRegexp(g, "")

}

func main() {
	demo()
	return

	re := `Hello, regular expression .+!`
	g, err := regstr.Parse(re, syntax.Perl)

	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		randstr, err := g.Gen()
		if err != nil {
			panic(err)
		}
		fmt.Println(randstr)
	}

	js := `\{".+": ".+"\}`
	g, err = regstr.Parse(js, syntax.Perl)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		randstr, err := g.Gen()
		if err != nil {
			panic(err)
		}
		fmt.Println(randstr)
	}

	sql := `INSERT INTO some_table\(col1,col2\) values \('.+','.+'\)`
	g, err = regstr.Parse(sql, syntax.Perl)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10; i++ {
		randstr, err := g.Gen()
		if err != nil {
			panic(err)
		}
		fmt.Println(randstr)
	}

}
