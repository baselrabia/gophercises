package bolt

import (
	"errors"
	bolt "go.etcd.io/bbolt"
)

var DB *bolt.DB

var (
	dbFileName      = "tasks.db"
	tasksBucketName = []byte("tasks")
)

func Connect() (*bolt.DB, error) {
	db, err := bolt.Open(dbFileName, 0600, nil)
	if err != nil {
		return nil, err
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(tasksBucketName)
		return err
	}); err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}

func Connection() (*bolt.DB, error) {
	if DB == nil {
		return nil, errors.New("there is no db connection")
	}
	return DB, nil
}
