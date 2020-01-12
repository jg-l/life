package main

import (
	"time"
)

const lifeLogFormat = "01/02 2006 15:04PM"

// Entry is the object that holds an entry
type Entry struct {
	Message      string `json:"message"`
	CreationDate int64  `json:"timestamp"`
}

func newEntry(message string) Entry {
	return Entry{message, time.Now().Unix()}
}

func (e Entry) getTimeObject() time.Time {
	return time.Unix(e.CreationDate, 0)
}

type Entries []Entry

func (e Entries) Len() int           { return len(e) }
func (e Entries) Less(i, j int) bool { return e[i].CreationDate > e[j].CreationDate }
func (e Entries) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
