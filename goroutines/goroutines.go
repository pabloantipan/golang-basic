package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Say(text string) string {
	return fmt.Sprintf("I say %s", text)
}

func SayIt(text string) {
	fmt.Println(Say(text))
}

func Worker(id int) {
	workerTaskTime := 1
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second * time.Duration(workerTaskTime))
	fmt.Printf("Worker %d done\n", id)
}

func DoWork(id int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	Worker(id)
}
