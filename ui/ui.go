package ui

import (
	"bufio"
	"fmt"
	"go-basic/channeling"
	"go-basic/goroutines"
	"go-basic/interfacing"
	"go-basic/mapping"
	"go-basic/palindrome"
	"go-basic/slicing"
	"go-basic/structs"
	"go-basic/utils"
	"os"
	"sync"
	"time"
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

func ShowStruct() {
	example := structs.Structing()
	fmt.Println("Show struct exmaple:", example)

	var exampleTwo structs.Car
	exampleTwo.Brand = "Ferrari"
	fmt.Println("Other example:", exampleTwo)
}

func ShowArray() {
	array := slicing.Arraying()
	fmt.Println(array, len(array), cap(array))
}

func ShowSlice() {
	slice := slicing.Slicing()
	fmt.Println(slice[0], slice[:3], slice[2:4], slice[4:])

	slice = append(slice, 7)
	fmt.Println(slice)

	newSlice := []int{11, 12, 13}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

func ShowPointers() {
	a := 50
	b := &a
	c := *b
	fmt.Println(a, b, c, &b)

	*b = 100
	fmt.Println(a)

	myPc := structs.Pc{
		Ram:   16,
		Disk:  200,
		Brand: "Acer",
	}
	fmt.Println(myPc)
	fmt.Println(myPc.PingPc())

	fmt.Println(myPc)
	myPc.DuplicateRAM()
	fmt.Println(myPc)
	myPc.DuplicateRAM()
	fmt.Println(myPc)
	myPc.DuplicateRAM()
	fmt.Println(myPc)
}

func ShowInterfacing() {
	aSquare := interfacing.Square{Base: 2}
	aRectangle := interfacing.Rectangle{Base: 3, Height: 2}

	fmt.Println("Area :", interfacing.Calculate(aSquare))
	fmt.Println("Area :", interfacing.Calculate(aRectangle))

	fmt.Println(interfacing.ShowInterfaceList()...)
}

func ShowGoRoutines() {
	goroutines.SayIt("Hey")
	go goroutines.SayIt("There!")
	time.Sleep(time.Second * 1)
}

func ShowGoRoutinesTwo() {
	var waitGroup sync.WaitGroup

	for i := 1; i < 5; i++ {
		waitGroup.Add(1)
		i := i
		go goroutines.DoWork(i, &waitGroup)
	}
	waitGroup.Wait()
}

func ShowGoRoutinesThree() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go goroutines.DoWork(1, &waitGroup)
	go goroutines.DoWork(2, &waitGroup)
	waitGroup.Wait()
}

func ShowChanneling() {
	channel := channeling.Channeling()
	fmt.Println(len(channel), cap(channel))
}
