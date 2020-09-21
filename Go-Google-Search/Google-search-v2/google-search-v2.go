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

func googleSearch(query string) (results []Result) {
	ch := make(chan Result)
	defer close(ch)
	go func() { ch <- web(query) }()
	go func() { ch <- image(query) }()
	go func() { ch <- video(query) }()
	for i := 0; i < 3; i++ {
		results = append(results, <-ch)
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
