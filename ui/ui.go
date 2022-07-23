package ui

import (
	"bufio"
	"fmt"
	"go-basic/palindrome"
	"go-basic/utils"
	"os"
)

func AskForPalindromeWord() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please give me palindrome word:")
	word, _ := reader.ReadString('\n')
	word = utils.CleaningWord(word)

	fmt.Println("You have entered word:", word)

	if palindrome.IsPalindromeShorter(word) {
		fmt.Println("Yep, it's a palindrome")
		return
	}

	fmt.Println("That's not a palindrome!")
}
