package backtracking

import (
	"slices"
	"strings"
)

func BackTracking(word string, dict []string, alph []string, answer *string) {

	if !strings.Contains(word, "*") {
		if slices.Contains(dict, word) {
			*answer = word
		}
		return
	}

	index := strings.Index(word, "*")

	for i := 0; i < len(alph); i++ {
		BackTracking(word[:index] + alph[i] + word[index + 1:], dict, alph, answer)
	}
}

