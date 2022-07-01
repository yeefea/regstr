package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp/syntax"
	"time"

	"github.com/spf13/cobra"
	"github.com/yeefea/regstr/gen"
)

var (
	repeatN int
	limit   int

	rootCmd = &cobra.Command{
		Use:   "regstr",
		Short: "regstr 'regular expression'",
		Args:  cobra.ExactArgs(1),
		RunE:  runE,
	}
)

func init() {
	rootCmd.Flags().IntVarP(&repeatN, "ntimes", "n", 1, "repeat n times")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 10, "limit")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Compile(re string) (*syntax.Regexp, error) {
	return syntax.Parse(re, syntax.Perl)
}

func runE(cmd *cobra.Command, args []string) error {
	if repeatN < 1 {
		return fmt.Errorf("n must be greater than 0")

	}
	if limit < 0 {
		return fmt.Errorf("limit must be greater than 0")
	}
	reg, err := Compile(args[0])
	if err != nil {
		return err
	}
	f := gen.StringGeneratorFactory{Limit: limit}
	g, err := f.NewStringGenerator(reg)
	if err != nil {
		return err
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < repeatN; i++ {
		s, err := g.Gen()
		if err != nil {
			return err
		}
		fmt.Println(s)
	}
	return nil
}
