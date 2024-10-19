package main

func firstOccurrence(input string) string {
	builder := ""

	seen := make(map[rune]bool)
	for _, ch := range input {
		if !seen[ch] {
			seen[ch] = true
			builder += string(ch)
		}
	}

	return builder
}
