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
	re := "Hello, regular expression .+\\!"
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

}
