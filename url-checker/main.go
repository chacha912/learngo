package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
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

	results := map[string]string{}
	c := make(chan requestResult)

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

// send only
func hitURL(url string, c chan<- requestResult) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK "
	if err != nil {
		status = "FAILED" + err.Error()
	} else if resp.StatusCode >= 400 {
		status = "FAILED" + resp.Status
	}
	c <- requestResult{url: url, status: status}
}
