package backtracking_test

import (
	"Worder/internal/utils/backtracking"
	"testing"
	"github.com/stretchr/testify/assert"
)






func TestBackTracking(t *testing.T) {
	tests := []struct {
		name         string
		word         string
		dictionary   []string
		alphabet     []string
		expectedWord string
	}{
		{
			name: "Single wildcard with valid replacement",
			word: "abab*abab",
			dictionary: []string{
				"ababaabab",
				"aabbbbbbbbbb",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "ababaabab",
		},
		{
			name: "Single wildcard with no valid replacement",
			word: "abab*abab",
			dictionary: []string{
				"xyzxyzxyz",
				"aabbbbbbbbbb",
			},
			alphabet:     []string{"x", "y", "z"},
			expectedWord: "",
		},
		{
			name: "Multiple wildcards with valid replacements",
			word: "ab*ab*ab*",
			dictionary: []string{
				"ababababab",
				"abbabbabb",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "abbabbabb",
		},
		{
			name: "All wildcards, valid word found",
			word: "********",
			dictionary: []string{
				"bbbbbbbb",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "bbbbbbbb",
		},
		{
			name: "All wildcards, no valid word",
			word: "********",
			dictionary: []string{
				"cccccccc",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "",
		},
		{
			name: "No wildcard, valid word in dictionary",
			word: "ababaabab",
			dictionary: []string{
				"ababaabab",
				"xyzxyzxyz",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "ababaabab",
		},
		{
			name: "No wildcard, word not in dictionary",
			word: "ababaabab",
			dictionary: []string{
				"xyzxyzxyz",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "",
		},
		{
			name: "Wildcard at start of word",
			word: "*bababab",
			dictionary: []string{
				"abababab",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "abababab",
		},
		{
			name: "Wildcard at end of word",
			word: "abababab*",
			dictionary: []string{
				"ababababa",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "ababababa",
		},
		{
			name: "Mixed wildcards",
			word: "*a*b*",
			dictionary: []string{
				"cabba",
				"babab",
			},
			alphabet:     []string{"a", "b", "c"},
			expectedWord: "cabba",
		},
		{
			name: "Empty dictionary",
			word: "abab*abab",
			dictionary:   []string{},
			alphabet:     []string{"a", "b"},
			expectedWord: "",
		},
		{
			name: "Empty alphabet",
			word: "abab*abab",
			dictionary: []string{
				"ababaabab",
			},
			alphabet:     []string{},
			expectedWord: "",
		},
		{
			name: "Empty word",
			word:         "",
			dictionary:   []string{"ababaabab"},
			alphabet:     []string{"a", "b"},
			expectedWord: "",
		},
		{
			name: "Wildcard with multiple valid replacements, first found",
			word: "abab*abab",
			dictionary: []string{
				"ababaabab",
				"ababababab",
			},
			alphabet:     []string{"a", "b"},
			expectedWord: "ababaabab",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var answer string
			backtracking.BackTracking(test.word, test.dictionary, test.alphabet, &answer)

			assert.Equal(t, test.expectedWord, answer, "Test case failed: %s", test.name)
		})
	}
}
