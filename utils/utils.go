package utils

import (
	"fmt"
	"strings"
)

func CleaningWord(word string) string {
	word = strings.ToLower(word)
	word = strings.Trim(word, "\r\n")
	return word
}

func ReverseWord(word string) string {
	word = strings.Trim(word, "\r\n")
	fmt.Println(word)
	byte_word := []rune(word)
	for i, j := 0, len(byte_word)-1; i < j; i, j = i+1, j-1 {
		byte_word[i], byte_word[j] = byte_word[j], byte_word[i]
	}
	fmt.Println(string(byte_word))
	return string(byte_word)
}
