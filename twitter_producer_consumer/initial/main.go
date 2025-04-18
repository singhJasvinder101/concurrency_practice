package main

import (
	"fmt"
	"time"
)

func produceTweets(stream *Stream) (tweets []*Tweet) {
	for {
		tweet, err := stream.Next()
		if err == errorEOF {
			break
		}

		if err != nil {
			fmt.Println("Unexpected error:", err)
			break
		}

		tweets = append(tweets, tweet)
	}
	return tweets
}

func consumer(tweets []*Tweet) {
	for _, t := range tweets {
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


	// Producer
	tweets := produceTweets(&stream)

	// Consumer
	consumer(tweets)
	fmt.Println("Execution time: ", time.Since(start))
}


