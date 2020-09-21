package main

import (
	"fmt"
	"math/rand"
	"time"
)

//variables storing functions returned as closure
var (
	web   = fakeSearch("web")
	image = fakeSearch("image")
	video = fakeSearch("video")
)

//Result stores the result in string format (anonymous literal)
type Result struct {
	string
}

func fakeSearch(query string) func(string) Result {
	return func(keyword string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("%s result for %s", query, keyword)}
	}
}

//Recover Function to handle panics when server takes more than specified time to return result
func Recover(keyword string) {
	if r := recover(); r != nil {
		fmt.Printf("Got Error in fetching %q results : %v", keyword, r)
	}
}

func googleSearch(query string) (results []Result) {
	ch := make(chan Result)

	go func() {
		defer Recover("web")
		ch <- web(query)
	}()

	go func() {
		defer Recover("image")
		ch <- image(query)
	}()

	go func() {
		defer Recover("video")
		ch <- video(query)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-ch:
			results = append(results, result)

		case <-timeout:
			fmt.Println("Request timed out...")
			close(ch)
			return
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	startTime := time.Now()
	query := "AWS"
	searchResult := googleSearch(query)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Showing search results for : %q\n\n", query)
	for _, result := range searchResult {
		fmt.Println("=> ", result.string)
	}
	fmt.Println("\nTime taken : ", elapsedTime)

}
