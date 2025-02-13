import (
	"fmt"
	"strings"
)

var set = make(map[string]struct{})

func setup(words []string) {
	for _, word := range words {
		for _, permutation := range findWildcardPermutations(word) {
			set[permutation] = struct{}{}
		}
	}
}

func findWildcardPermutations(word string) []string {
	var permutations []string
	for i := 0; i < len(word); i++ {
		permutations = append(permutations, permutate(word, i))
	}
	return permutations
}

func permutate(word string, i int) string {
	var builder strings.Builder
	for k := 0; k < len(word); k++ {
		if i != k {
			builder.WriteByte(word[k])
		} else {
			builder.WriteByte('*')
		}
	}
	return builder.String()
}

func hasWord(word string) bool {
	_, exists := set[word]
	return exists
}
