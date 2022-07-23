package palindrome

import (
	"go-basic/utils"
)

func IsPalindrome(word string) bool {
	word = utils.CleaningWord(word)
	byte_word := []rune(word)
	for i, j := 0, len(byte_word)-1; i < j; i, j = i+1, j-1 {
		if byte_word[i] != byte_word[j] {
			return false
		}
		byte_word[i], byte_word[j] = byte_word[j], byte_word[i]
	}
	return true
}

func IsPalindromeShorter(word string) bool {
	word = utils.CleaningWord(word)
	byte_word := []rune(word)
	for i := 0; i < len(byte_word)-1; i += 1 {
		if byte_word[i] != byte_word[len(byte_word)-1-i] {
			return false
		}
	}
	return true
}
