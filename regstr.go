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

const (
	regMobile = `13[0-9]{9}`
	regEmail  = `[a-z0-9]+@[a-z]+\.com`
	regInt    = `-?[1-9][0-9]*`
	regFloat  = `-?[1-9][0-9]*\.[0-9]+`
	regDate   = `(19|20)\d{2}-(0[1-9]|1[0-2])-((0[1-9])|(1[0-9]|2[0-8]))`
	regText   = `[A-Z][a-z]{0,6}( [a-z]{1,7}){5,15}\.( [A-Z][a-z]{0,6}( [a-z]{1,7}){5,15}\.)*`
)

var (
	repeatN int
	limit   int
	in      string
	out     string

	rootCmd  = &cobra.Command{Use: "regstr"}
	regexCmd = &cobra.Command{
		Use:   "regex [regular expression]",
		Short: "Generate random strings from the given regular expression.",
		Args:  cobra.MaximumNArgs(1),
		RunE:  runReg,
	}
	mobileCmd = &cobra.Command{
		Use:   "mobile",
		Short: "Generate random mobile numbers.",
		Args:  cobra.NoArgs,
		RunE:  runSimple(regMobile),
	}
	emailCmd = &cobra.Command{
		Use:   "email",
		Short: "Generate random E-Mails.",
		Args:  cobra.NoArgs,
		RunE:  runSimple(regEmail),
	}
	intCmd = &cobra.Command{
		Use:   "int",
		Short: "Generate random integer numbers.",
		Args:  cobra.NoArgs,
		RunE:  runSimple(regInt),
	}
	floatCmd = &cobra.Command{
		Use:   "float",
		Short: "Generate random floating point numbers.",
		Args:  cobra.NoArgs,
		RunE:  runSimple(regFloat),
	}
	dateCmd = &cobra.Command{
		Use:   "date",
		Short: "Generate random date strings.",
		Args:  cobra.NoArgs,
		RunE:  runSimple(regDate),
	}
	textCmd = &cobra.Command{
		Use:   "text",
		Short: "Generate random text.",
		Args:  cobra.NoArgs,
		RunE:  runSimple(regText),
	}
)

func init() {
	regexCmd.Flags().StringVarP(&in, "input", "i", "", "input file")

	rootCmd.PersistentFlags().StringVarP(&out, "output", "o", "", "output file")
	rootCmd.PersistentFlags().IntVarP(&repeatN, "ntimes", "n", 1, "repeat n times")
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "l", 10, "limit")
	rootCmd.AddCommand(regexCmd)
	rootCmd.AddCommand(mobileCmd, emailCmd, intCmd, floatCmd, dateCmd, textCmd)

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

func validateParams() error {
	if repeatN < 1 {
		return fmt.Errorf("n must be greater than 0")

	}
	if limit < 0 {
		return fmt.Errorf("limit must be greater than 0")
	}
	return nil
}

func runReg(cmd *cobra.Command, args []string) error {
	if err := validateParams(); err != nil {
		return err
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

	return runGeneral(re)

}

func runGeneral(re string) error {
	if err := validateParams(); err != nil {
		return err
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

func runSimple(re string) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error { return runGeneral(re) }
}
