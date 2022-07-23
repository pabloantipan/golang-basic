package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// var array [4]int
	// fmt.Println(array)

	// slice := []int{0, 1, 2, 3, 4, 5, 6}
	// fmt.Println(slice, cap(slice))

	// message := greetings.Hello("Gladys")
	// fmt.Println(message)

	// slicer := []string{"How", "are", "you", "doing", "?"}

	// for i, value := range slicer {
	// 	fmt.Println(i, value)
	// }

	askForPalindromeWord()
}

func reverseWord(word string) string {
	word = strings.Trim(word, "\r\n")
	fmt.Println(word)
	byte_word := []rune(word)
	for i, j := 0, len(byte_word)-1; i < j; i, j = i+1, j-1 {
		byte_word[i], byte_word[j] = byte_word[j], byte_word[i]
	}
	fmt.Println(string(byte_word))
	return string(byte_word)
}

func isPalindrome(word string) bool {
	word = strings.Trim(word, "\r\n")
	byte_word := []rune(word)
	for i, j := 0, len(byte_word)-1; i < j; i, j = i+1, j-1 {
		if byte_word[i] != byte_word[j] {
			return false
		}
		byte_word[i], byte_word[j] = byte_word[j], byte_word[i]
	}
	return true
}

func isPalindromeShorter(word string) bool {
	word = strings.ToLower(strings.Trim(word, "\r\n"))
	byte_word := []rune(word)
	for i := 0; i < len(byte_word)-1; i += 1 {
		if byte_word[i] != byte_word[len(byte_word)-1-i] {
			return false
		}
	}
	return true
}

func askForPalindromeWord() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please give me palindrome word:")
	word, _ := reader.ReadString('\n')

	fmt.Println("You have entered word:", word)

	if isPalindromeShorter(word) {
		fmt.Println("Yep, it's a palindrome")
		return
	}

	fmt.Println("That's not a palindrome!")
}
