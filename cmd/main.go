package main

import (
	"fmt"
	"math/rand"
	"regexp/syntax"
	"time"
	"unicode"

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

	for i := 0; i < 1000; i++ {
		var x rune = rand.Int31n(unicode.MaxRune)
		fmt.Printf("%d %c\n", x, x)
	}

	re := `[^我]`
	g, err := syntax.Parse(re, syntax.Perl)
	if err != nil {
		panic(err)
	}
	DescRegexp(g, "")
	g.Simplify()
	DescRegexp(g, "")

}

func main() {

	re := `aaa|bbb`
	g := regstr.MustCompile(re)

	for i := 0; i < 10; i++ {
		randstr, err := g.Gen()
		if err != nil {
			panic(err)
		}
		fmt.Println(randstr)
	}

	js := `\{".+": ".+"\}`
	g, err := regstr.Compile(js)
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

	return

	sql := `INSERT INTO \nsome_table\(col1,col2\) values \n\('.+','.+'\)(,\n\('.+','.+'\))+;`
	g, err = regstr.Compile(sql)
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
