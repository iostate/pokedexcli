package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	internal "github.com/iostate/pokedexcli/internal/pokeapi"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := lowercaseAndBreakUpWords(c.input)

		// Check length of actual against length of expected
		if len(actual) != len(c.expected) {
			t.Errorf("length of actual not equal to length of expected, %d != %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("got %s want %s", word, expectedWord)
			}
		}
	}
}

func TestCommands(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create a config instance to pass to commandHelp
	cfg := &config{
		client: internal.NewClient(5*time.Second, 5*time.Second),
	}

	err := commandHelp(cfg)
	if err != nil {
		t.Errorf("commandHelp returned an error: %v", err)
	}

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "help:") {
		t.Errorf("help command output missing 'help:' command")
	}
	if !strings.Contains(output, "exit:") {
		t.Errorf("help command output missing 'exit:' command")
	}
}
