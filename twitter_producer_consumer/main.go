package main

import (
	"fmt"
	"sync"
	"time"
)

func produceTweets(stream *Stream, tweetChan chan *Tweet){
	for {
		tweet, err := stream.Next()
		if err != nil {
			if err != errorEOF {
				fmt.Println("stop when there's an EOF or real error")
			}
			fmt.Println("Error reading stream:", err)
			break 
		}

		
		tweetChan <- tweet
		// tweets = append(tweets, tweet)
	}
}

func consumer(tweetChan chan *Tweet, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// t := <- tweetChan
	for t := range tweetChan{
		if t.IsTalkingAboutGO() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}


func main (){
	start := time.Now()
	stream := GetMockStream()

	var tweetChan = make(chan *Tweet)
	// var tweets []*Tweet
	
	var wg sync.WaitGroup
	wg.Add(2)
	
	// Producer
	go func(){
		defer wg.Done()	
		
		// tweets = produceTweets(&stream, tweetChan)
		produceTweets(&stream, tweetChan)
		close(tweetChan)
	}()

	// Consumer
	go consumer(tweetChan, &wg)

	wg.Wait()
	fmt.Println("Execution time: ", time.Since(start))
}


