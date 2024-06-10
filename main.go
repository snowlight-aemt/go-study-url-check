package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func main() {
	hitUrl()
}

func hitUrl() {
	var results = map[string]string{}
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
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}

		results[url] = result
	}

	for key, value := range results {
		fmt.Println(key, value)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking : ", url)
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println("Err : ", url, resp.StatusCode)
		return errRequestFailed
	}
	return nil
}
