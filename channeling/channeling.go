package channeling

func Channeling() (int, int) {
	channel := make(chan string, 2)
	channel <- "Message 1"
	channel <- "Message 2"

	// Range and close
	close(channel)

	return len(channel), cap(channel)
}
