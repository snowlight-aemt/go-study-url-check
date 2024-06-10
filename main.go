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
		"https://www.google.com/",
		"https://www.airbnb.co.kr/",
		"https://aws.amazon.com/",
		"https://www.facebook.com/",
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
	req, err := http.NewRequest("GET", url, nil)
	//req.Header.Add("Accept-Encoding", "gzip, deflate")
	//req.Header.Add("Accept", "*/*"
	//req.Header.Add("Accept-Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)

	//resp, err := http.Get(url)

	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Body)

		status = "FAILED"
	}

	result <- httpResult{url, status}
}
