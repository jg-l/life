package main

import (
	"os"
	"sort"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {

	e := Entries{}
	e.load() // Load the database

	a := newEntry("This is my test entry")

	err := e.save(a)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if len(e) != 1 {
		t.Log("Got:", len(e), "Want:", 1)
		t.Fail()
	}

	os.RemoveAll(getLifeLogLocation())
}

func TestSortTasks(t *testing.T) {

	entries := Entries{}
	entries.load()

	last := newEntry("This is should be last")
	time.Sleep(3 * time.Second)
	entries.save(last)

	middle := newEntry("This is should be in the middle")
	time.Sleep(2 * time.Second)
	entries.save(middle)

	first := newEntry("This should be first")
	time.Sleep(1 * time.Second)
	entries.save(first)
	sort.Sort(entries)
	counter := 0
	for e := range entries {
		switch counter {
		case 0:
			if entries[e].Message != first.Message {
				t.Log(entries[e].Message, first.Message)
				t.Fail()
			}
		case 1:
			if entries[e].Message != middle.Message {
				t.Log(entries[e].Message, middle.Message)
				t.Fail()
			}
		case 2:
			if entries[e].Message != last.Message {
				t.Log(entries[e].Message, last.Message)
				t.Fail()
			}

		}
		counter++
	}

	os.RemoveAll(getLifeLogLocation())
}
