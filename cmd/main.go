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

func main() {
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
