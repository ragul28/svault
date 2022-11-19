package vault

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/boltdb/bolt"
)

func open(file string) *bolt.DB {
	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	db, err := bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	return db
}

func writeDB(db *bolt.DB, bucket, key string, value []byte) error {
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), []byte(value))
		return err
	})
	return nil
}

func readDB(db *bolt.DB, bucket, key string) (value []byte, errres error) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %q not found", bucket)
		}

		value = b.Get([]byte(key))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func deleteDB(db *bolt.DB, bucket, key string) error {
	err := db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucket))
		b.Delete([]byte(key))

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func iterateDB(db *bolt.DB, bucket string) (counter int, err error) {

	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("svault is empty!")
			// fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
		}
	}()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		c := b.Cursor()
		counter = 0
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			counter++
			fmt.Printf("%d. %s\n", counter, k)
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return 0, nil
}
