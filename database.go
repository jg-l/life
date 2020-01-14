package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func startOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func getLifeLogLocation() string {
	u, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return u + ".life"
}

// save is simple but expensive, but its for single user case, so performance isn't the main goal.
// loads entries first
// Append new one
// save Entries into the file
func (e *Entries) save(ne Entry) error {
	db, err := os.OpenFile(getLifeLogLocation(), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	*e = append(*e, ne)
	b, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		return err
	}

	_, err = db.Write(b)
	if err != nil {
		return err
	}

	return nil

}

func (e *Entries) load() {

	f, err := os.OpenFile(getLifeLogLocation(), os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	db, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(db, e)

}
