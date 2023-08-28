package helpers

import (
	"strings"
	"unicode"
)

func CountChars(s string) (int, int) {
	var cons int
	var vow int
	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) {
			vowels := "aeiou"
			if strings.ContainsRune(vowels, c) {
				vow += 1
			} else {
				cons += 1
			}
		}
	}
	return cons, vow
}

func ShareCommonFactors(a, b int) bool {
	cf := gcd(a, b)
	return cf > 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
