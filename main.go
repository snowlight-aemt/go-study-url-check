package main

import (
	"fmt"
	"net/http"
)

type httpResult struct {
	url    string
	status string
}

func main() {
	results := map[string]string{}
	c := make(chan httpResult)
	urls := []string{
		"https://www.youtube.com/",
		"https://fastcampus.co.kr/",
		"https://www.inflearn.com/",
		"https://www.google.com/",
		"https://www.reddit.com/",
		//"https://www.airbnb.co.kr/",
		"https://aws.amazon.com/",
		"https://yomou.syosetu.com/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for key, val := range results {
		fmt.Println(key, val)
	}
}

func hitURL(url string, result chan<- httpResult) {
	resp, err := http.Get(url)

	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	result <- httpResult{url, status}
}
