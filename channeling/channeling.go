package channeling

type EmailChannels struct {
	EmailOne chan string
	EmailTwo chan string
}

func (emailChannels *EmailChannels) Channeling() {
	emailChannels.EmailOne = make(chan string)
	emailChannels.EmailTwo = make(chan string)
}

func message(text string, channel chan<- string) {
	channel <- text
}

func (emailChannels *EmailChannels) MessageOnOne(text string) {
	message(text, emailChannels.EmailOne)
}

func (emailChannels *EmailChannels) MessageOnTwo(text string) {
	message(text, emailChannels.EmailOne)
}

func (emailChannels *EmailChannels) CloseBoth() {
	close(emailChannels.EmailOne)
	close(emailChannels.EmailTwo)
}
