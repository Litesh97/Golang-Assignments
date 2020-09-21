package main

import (
	"fmt"
	"math/rand"
	"time"
)

//variables storing functions returned as closure
var (
	web    = fakeSearch("web")
	image  = fakeSearch("image")
	video  = fakeSearch("video")
	web2   = fakeSearch("web")
	image2 = fakeSearch("image")
	video2 = fakeSearch("video")
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

//Recover Function to handle panics when server takes more than specified time to return result and
//attemp is done to write on closed channels
func Recover(keyword interface{}) {
	if r := recover(); r != nil {
		switch keyword.(type) {
		case string:
			//log.Printf("Got Error in fetching %q results : %v\n", keyword, r)
		case int:
			//log.Printf("Replica %d was slow for result.Discarded!\n", keyword)
		}

	}
}

func replicate(query string, replicas ...func(string) Result) Result {
	ch := make(chan Result)

	searchForReplica := func(i int) {
		defer Recover(i)
		ch <- replicas[i](query)
	}
	for i := range replicas {
		go searchForReplica(i)
	}
	defer close(ch)
	return <-ch
}

func googleSearch(query string) (results []Result) {
	ch := make(chan Result)

	go func() {
		defer Recover("web")
		ch <- replicate(query, web, web2)
	}()

	go func() {
		defer Recover("image")
		ch <- replicate(query, image, image2)
	}()

	go func() {
		defer Recover("video")
		ch <- replicate(query, video, video2)
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
	query := "Druva"
	searchResult := googleSearch(query)
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nShowing search results for : %q\n\n", query)
	for _, result := range searchResult {
		fmt.Println("=> ", result.string)
	}
	fmt.Println("\nTime taken : ", elapsedTime)

}
