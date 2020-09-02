package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

// HEY ALEX
// customize the top-level variables below to tweak the linter.

// these are failures that match exactly, and are expressly disallowed.
var literalFailures = []string{
	"{{",
	"}}",
	"{%",
	"%}",
	"{#",
	"#}",
}

// this is the regex that tries to find lists. it's probably not perfect.
var listsRegex = regexp.MustCompile(`^\s*[a-zA-Z0-9]+\.\s*`)

// this detects github flavored markdown-style code snippets, to warn the user they won't work.
var codeRegex = regexp.MustCompile("^```[a-zA-Z0-9]*$")

func main() {
	app := cli.NewApp()
	app.Version = "0.1.0"

	app.Action = lint

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func lint(ctx *cli.Context) error {
	if len(ctx.Args()) != 1 {
		return errors.New("filename required (pass /dev/stdin if you want to pipe)")
	}

	f, err := os.Open(ctx.Args()[0])
	if err != nil {
		return fmt.Errorf("trouble opening file: %w", err)
	}

	var line int

	s := bufio.NewScanner(f)
	for s.Scan() {
		line++
		if err := lineLint(s.Text()); err != nil {
			return fmt.Errorf("line %d has errors: %w", line, err)
		}
	}

	fmt.Println("no errors!")

	return nil
}

func lineLint(s string) error {
	for _, f := range literalFailures {
		if strings.Contains(s, f) {
			return fmt.Errorf("%q found in line", f)
		}
	}

	if listsRegex.MatchString(s) {
		return errors.New("ordered list detected in line")
	}

	if codeRegex.MatchString(s) {
		return errors.New("GFM-style code snippet found in line")
	}

	return nil
}
