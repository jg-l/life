package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	args := os.Args[1:]

	entries := Entries{}
	entries.load()

	if len(args) == 0 {
		if len(entries) == 0 {
			fmt.Println("There are no life entries")
			os.Exit(2)
		}
		sort.Reverse(entries)
		display(entries)
	} else {
		message := strings.Join(args[:], " ")
		entry := newEntry(message)
		entries.save(entry)
	}

}
