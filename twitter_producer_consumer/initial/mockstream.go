package main

import (
	"errors"
	"strings"
	"time"
)

type Tweet struct {
	Username string
	Text     string
}

var mockdata = []Tweet{
	{
		"davecheney",
		"#golang top tip: if your unit tests import any other package you wrote, including themselves, they're not unit tests.",
	}, {
		"beertocode",
		"Backend developer, doing frontend featuring the eternal struggle of centering something. #coding",
	}, {
		"ironzeb",
		"Re: Popularity of Golang in China: My thinking nowadays is that it had a lot to do with this book and author https://github.com/astaxie/build-web-application-with-golang",
	}, {
		"beertocode",
		"Looking forward to the #gopher meetup in Hsinchu tonight with @ironzeb!",
	}, {
		"vampirewalk666",
		"I just wrote a golang slack bot! It reports the state of github repository. #Slack #golang",
	},
}


type Stream struct {
	pos    int
	tweets []Tweet
}

var errorEOF = errors.New("EOF: End of file")

// Next returns the next Tweet in the stream, returns EOF error if
// there are no more tweets
func (s *Stream) Next() (*Tweet, error) {
	// simulate delay 
	time.Sleep(500 * time.Millisecond)
	if s.pos >= len(s.tweets){
		return &Tweet{}, errorEOF
	}

	tweet := s.tweets[s.pos]
	s.pos++

	return &tweet, nil
}

func (t *Tweet) IsTalkingAboutGO() bool {
	// simulate delay
	time.Sleep(500 * time.Millisecond)

	hasGo := strings.Contains(strings.ToLower(t.Text), "go")
	hasGopher := strings.Contains(strings.ToLower(t.Text), "gopher")

	return hasGo || hasGopher
}

func GetMockStream() Stream {
	return Stream{0, mockdata}
}
