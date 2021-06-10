package main

import (
	"fmt"
	"time"

	"jaytaylor.com/archive.today"
)

var searchURL = "https://jaytaylor.com/"

func main() {
	snapshots, err := archivetoday.Search(searchURL, 10*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", snapshots)
}
