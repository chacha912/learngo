package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	// results := map[string]string{}
	c := make(chan result)

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}

// send only
func hitURL(url string, c chan<- result) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK "
	if err != nil {
		status = "FAILED" + err.Error()
	} else if resp.StatusCode >= 400 {
		status = "FAILED" + resp.Status
	}
	c <- result{url: url, status: status}
}
