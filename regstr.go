package main

import (
	"fmt"
	"io/ioutil"
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
	in      string
	out     string
	rootCmd = &cobra.Command{
		Use:   "regstr",
		Short: "regstr \"regex\"",
		Args:  cobra.MaximumNArgs(1),
		RunE:  runE,
	}
)

func init() {
	rootCmd.Flags().IntVarP(&repeatN, "ntimes", "n", 1, "repeat n times")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 10, "limit")
	rootCmd.Flags().StringVarP(&in, "input", "i", "", "input file")
	rootCmd.Flags().StringVarP(&out, "output", "o", "", "output file")
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
	if in == "" && len(args) == 0 {
		return fmt.Errorf("either an input file or a regular expression should be specified")
	}

	var re string
	if in != "" {
		data, err := ioutil.ReadFile(in)
		if err != nil {
			return err
		}

		re = string(data)

	} else {
		re = args[0]
	}
	reg, err := Compile(re)
	if err != nil {
		return err
	}
	f := gen.StringGeneratorFactory{Limit: limit}
	g, err := f.NewStringGenerator(reg)
	if err != nil {
		return err
	}

	rand.Seed(int64(time.Now().Nanosecond()))

	var outFil *os.File
	if out == "" {
		outFil = os.Stdout
	} else {
		outFil, err = os.Create(out)
		if err != nil {
			return err
		}
	}
	for i := 0; i < repeatN; i++ {
		s, err := g.Gen()
		if err != nil {
			return err
		}
		fmt.Fprintln(outFil, s)
	}
	return nil
}
