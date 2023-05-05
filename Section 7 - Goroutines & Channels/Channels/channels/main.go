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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// How you create a channel. Simple.
	c := make(chan string)

	// First loop to make initial requests
	for _, link := range links {
		go checkLink(link, c)
	}

	// Instead of repeatedly extracting from the channel to retrieve
	// our results using consecutive <-c statements like this...
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	
	// Second loop to continually make requests
	// ...We can use a for loop with the channel instead:
	//for i := 0; i < len(links); i++ {
	//for {  // Produces an infinite loop
	for l := range c {  // This syntax is equivalent to "for {", except it's more descriptive

		// PROTIP: When using range with a channel, we're telling the runtime
		//		   environment to "wait for the channel to return some value.
		//		   After it has returned some value, assign it to the variable l
		//		   and run the body of the for loop".

		// Note that <-c can block the for loop from continuing on
		//fmt.Println(<-c)
		//go checkLink(<-c, c)
		//go checkLink(l, c)

		// Use a function literal to sleep a few seconds between each get request. We
		// didn't want to put this sleep inside of checkLink because that would have
		// caused a sleep during the inital call to checklink above as well!
		go func(link string) {
			time.Sleep(5 * time.Second)
			// Don't just pass l from the outer scope to checkLink. The code will compile and run,
			// but you'll get a warning like: "loop variable l captured by func literalloopclosure"
			// Basically, sharing variables like this b/w different gorotuines (i.e., b/w the main
			// goroutine and its child goroutines) will cause weird stuff to happen, especially if
			// we end up changing that variable. Instead, make sure to use a copy of l that is passed
			// to this function literal.
			checkLink(link, c)
		}(l)
	}
}


func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println(link, " is up!")
	c <- link
}