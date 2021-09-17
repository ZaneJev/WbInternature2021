package main

import "time"

func ownSleep(t time.Duration) {
	<-time.After(time.Second*t)
}

func main() {
	ownSleep(10)
}