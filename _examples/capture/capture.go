package main

import (
	"fmt"

	"jaytaylor.com/archive.today"
)

var captureURL = "https://jaytaylor.com/"

func main() {
	archiveURL, err := archivetoday.Capture(captureURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully archived %v via archive.is: %v\n", captureURL, archiveURL)
}
