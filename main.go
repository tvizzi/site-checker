package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://github.com",
		"http://stackoverflow.com",
		"https://ggsel.net",
		"https://dzen.ru",
		"https://fedoraproject.org",
		"https://www.mvideo.ru",
		"https://www.chess.com",
	}

	c := make(chan string)

	start := time.Now()

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

	fmt.Printf("\nTotal time taken: %s\n", time.Since(start))
}

func checkUrl(url string, c chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("[ERROR] %s is down!", url)
		return
	}
	defer resp.Body.Close()

	duration := time.Since(start)
	c <- fmt.Sprintf("[OK] %s Status: %d (%s)", url, resp.StatusCode, duration)
}
