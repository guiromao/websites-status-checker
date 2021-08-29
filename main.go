package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	/* Creating a channel to communicate
	in between go routines, including the main one */
	channel := make(chan string)

	for _, link := range links {
		//write "go" to create a new go-routine/thread
		//only before function calls
		go checkLink(link, channel)
	}

	//the new threads are created, so the program will move
	//to this 'for' below

	//infinite for loop
	for l := range channel {
		//this is a blocking code
		//as receiving messages blocks the thread
		// --> fmt.Println(<-channel)

		//the checkLink will send the link to the channel
		//we'll call again to do a continuous status check
		//plus, receiving the message is a blocking code, '<-channel'
		//go checkLink(l, channel)
		// same as go checkLink(<-channel, channel), if for didn't have range of channel

		go func(link string) {
			time.Sleep(time.Second * 2)
			checkLink(link, channel)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	result := ""
	_, err := http.Get(link)

	if err != nil {
		result = "might be down."
		//c <- "Might be down, I think..."
	} else {
		result = "is up!"
		//c <- "Yep, " + link + " is up!"
	}

	fmt.Println(link, result)

	c <- link
}
