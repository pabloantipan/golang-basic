package goroutines

import "fmt"

func Say(text string) string {
	return fmt.Sprintf("I say %s", text)
}

func SayIt(text string) {
	fmt.Println(Say(text))
}
