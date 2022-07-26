package ui

import (
	"bufio"
	"fmt"
	"go-basic/mapping"
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

func ShowMapping() {
	example := mapping.Mapping()
	fmt.Println("This is a map: ", example)

	// Running a map. WARNING: It's random. Use slice for get ordered
	fmt.Println("\nRunning the map:")
	for idx, value := range example {
		fmt.Println(idx, value)
	}

	// Find a value
	key := "Joseph"
	valueFound, isExist1 := example[key]
	fmt.Println("\nFound value for key=", key, ":", valueFound, isExist1)

	// Find wrong key
	wrongKey := "Josephh"
	wrongValueFound, isExist2 := example[wrongKey]
	fmt.Println("\nFound value for a wrong key=", key, ":", wrongValueFound, isExist2)
}
