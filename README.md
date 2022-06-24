# RegStr
Generate random strings from regular expressions.

## Examples

```go
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
```


The output will be:
```
Hello, regular expression 1rrTkBZIynysBWv!
Hello, regular expression uX4lCw4b!
Hello, regular expression MQLmA0uZy!
Hello, regular expression Oms4e83k!
Hello, regular expression bLOF9JQ!
Hello, regular expression BWr2rqOzy1a!
Hello, regular expression ojn8PJ8!
Hello, regular expression 9JtHOYEkGb!
Hello, regular expression LPU!
Hello, regular expression KtCjKtKze!
```