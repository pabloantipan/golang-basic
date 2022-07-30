package channeling

func Channeling() <-chan string {
	channel := make(chan string, 2)
	channel <- "Message 1"
	channel <- "Message 2"

	// Range and close
	close(channel)

	return channel
}
