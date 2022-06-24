package regstr

import (
	"fmt"
	"regexp"
	"regexp/syntax"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	regex := "test\\.asd+\\.qq\\.com[a-z]"
	re := regexp.MustCompile(regex)
	g, err := Parse(regex, syntax.Perl)
	assert.Nil(t, err)
	for i := 0; i < 10; i++ {
		randomStr, err := g.Gen()
		assert.Nil(t, err)
		fmt.Println(randomStr)
		assert.True(t, re.Match([]byte(randomStr)))
	}
}
