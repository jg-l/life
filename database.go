package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

const bucketName = "logs"

func startOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func getLifeLogLocation() string {
	u, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return u + "/.life"
}

func createEntriesBucket(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}

func initDb() {
	db, err := getDb()
	defer db.Close()

	if err != nil {
		panic(err)
	}

	err = createEntriesBucket(db)

	if err != nil {
		panic(err)
	}

}

func getDb() (*bolt.DB, error) {
	db, err := bolt.Open(getLifeLogLocation(), 0600, &bolt.Options{Timeout: 1 * time.Second})
	return db, err
}

func save(e Entry) {

	db, err := getDb()
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		taskJSON, err := json.Marshal(e)

		if err != nil {
			panic(err)
		}

		now := time.Now()
		key := now.Format(time.RFC3339)

		err = b.Put([]byte(key), taskJSON)
		return err
	})

}

func getEntries() Entries {
	db, err := getDb()
	defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var entries Entries

	// Get All Entries for now
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var entry Entry
			if err := json.Unmarshal([]byte(v), &entry); err != nil {
				log.Fatalln(err)
			}
			entries = append(entries, entry)

		}

		return nil

	})

	// Range Scan using date
	// db.View(func(tx *bolt.Tx) error {
	// 	c := tx.Bucket([]byte(bucketName)).Cursor()

	// 	now := time.Now()
	// 	nowKey := now.Format(time.RFC3339)
	// 	today := startOfDay(now)
	// 	// yesterday := today.AddDate(0, 0, -1)
	// 	todayKey := today.Format(time.RFC3339)
	// 	// yesterdayKey := yesterday.Format(time.RFC3339)

	// 	min := []byte(todayKey)
	// 	max := []byte(nowKey)

	// 	for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {

	// 		var e Entry
	// 		err := json.Unmarshal(v, &e)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		entries = append(entries, e)
	// 	}

	// 	return nil
	// })
	return entries
}
