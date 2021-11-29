package reverseWords

import (
	"fmt"
	"regexp"
	"strings"
)

func ReverseWords(s string) string {

	regex := regexp.MustCompile(`\s+`)
	tokens := regex.Split(strings.TrimSpace(s), -1)

	for j, token := range tokens {
		runes := []rune(token)
		for i := 0; i < len(runes)/2; i++ {
			runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
		}
		tokens[j] = string(runes)
	}

	fmt.Println(tokens)

	return strings.Join(tokens, " ")
}
