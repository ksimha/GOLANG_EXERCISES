package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		// go keyword is like a thread
		go checkLink(link, c)
	}

	for l := range c {
		time.Sleep(5 * time.Second)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
