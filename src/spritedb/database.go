package spritedb

import (
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)

type DB struct {
	currentDatabase string
	*Options
	db *bolt.DB // Underlying BoltDB instance
}

// NewDB function
func NewDB(opts ...OptFunc) (*DB, error) {
	opt := &Options{
		Encoder: &JSONEncoder{}, // Use pointer here
		Decoder: &JSONDecoder{}, // Use pointer here
		DBName:  "default",      // Provide a default DB name
	}

	for _, fn := range opts {
		fn(opt)
	}

	dbName := fmt.Sprintf("%s.db", opt.DBName) // Use opt.DBName
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("collection-meta"))
		return err
	}); err != nil {
		return nil, err
	}

	return &DB{
		currentDatabase: opt.DBName,
		db:              db,
		Options:         opt, // Pass the dereferenced Options struct
	}, nil
}

// DropDatabase removes the database file associated with the current database.
//
// It returns an error if there was a problem deleting the file.
func (db *DB) DropDatabase() error {
	dbName := fmt.Sprintf("%s.db", db.DBName)
	db.Close()
	return os.Remove(dbName)
}

func (db *DB) Close() {
	db.db.Close()
}
