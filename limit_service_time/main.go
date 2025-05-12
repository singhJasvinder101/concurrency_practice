//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync/atomic"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	}

	timer := make(chan bool, 1)
	remainingTime := 10 - atomic.LoadInt64(&u.TimeUsed)
	start := time.Now() // resets every time func called
	
	// timeout goroutine for free users
	go func(){
		process()
		timer <- true
	}()

	select {
		case <- timer:
			atomic.AddInt64(&u.TimeUsed, int64(time.Since(start).Seconds()))
			return true
		// case <- time.After(10 * time.Second):
		case <- time.After(time.Duration(remainingTime) * time.Second):
			atomic.AddInt64(&u.TimeUsed, int64(remainingTime))
			return false
	}
}

func main() {
	RunMockServer()
}