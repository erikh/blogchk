package main

import "testing"

func TestLiterals(t *testing.T) {
	for _, f := range literalFailures {
		if err := lineLint(f); err == nil {
			t.Fatal("literal failure check should not have passed")
		}
	}

	table := map[string]string{
		"braces with other text": "{{ other text",
		"braces at end of line":  "other text }}",
		"braces in middle":       "other {{ }} text",
	}

	for _, f := range table {
		if err := lineLint(f); err == nil {
			t.Fatal("literal failure check should not have passed")
		}
	}
}

func TestOrderedLists(t *testing.T) {
	table := map[string]string{
		"ordered list with integer":              "1. this is an ordered list",
		"ordered list with non-starting integer": "10. this is a really long ordered list",
		"ordered list with alpha chars":          "A. this is an ordered list",
		"ordered list with lowercase alpha":      "a. this is an ordered list",
		"ordered list with multi alpha":          "iii. this is the third argument in an ordered sublist",
	}

	for _, f := range table {
		if err := lineLint(f); err == nil {
			t.Fatal("literal failure check should not have passed")
		}
	}
}

func TestSnippets(t *testing.T) {
	table := map[string]string{
		"blank snippet": "```",
		"ruby snippet":  "```ruby",
		"bash snippet":  "```bash",
	}

	for _, f := range table {
		if err := lineLint(f); err == nil {
			t.Fatal("literal failure check should not have passed")
		}
	}
}
