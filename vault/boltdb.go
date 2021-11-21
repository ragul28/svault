package vault

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type dbops interface {
	writeDB(db *bolt.DB, bucket, key, value string) error
	readDB(db *bolt.DB, bucket, key string) (value string)
	deleteDB(db *bolt.DB, bucket, key string) error
	iterate(db *bolt.DB, bucket string) error
}

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

func readDB(db *bolt.DB, bucket, key string) (value string) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", bucket)
		}

		value = string(b.Get([]byte(key)))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return value
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

func iterate(db *bolt.DB, bucket string) error {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
