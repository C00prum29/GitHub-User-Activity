package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	username := os.Args[1]

	events, err := FetchEvents(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(events) == 0 {
		fmt.Println("No activity found.")
		return
	}

	for _, e := range events {
		fmt.Println(FormatEvent(e))
	}
}
