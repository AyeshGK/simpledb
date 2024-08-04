package spritedb

import bolt "go.etcd.io/bbolt"

type Collection struct {
	bucket *bolt.Bucket
}

func (db *DB) CreateCollection(name string) (*Collection, error) {
	collection := &Collection{
		bucket: nil,
	}
	err := db.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucket([]byte(name))
		if err != nil {
			return err
		}

		collection.bucket = bucket
		return nil
	})
	if err != nil {
		return nil, err
	}

	_, err = db.CreateCollectionMeta(name)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func (db *DB) DropCollection(collectionName string) error {
	err := db.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(collectionName))
	})
	if err != nil {
		return err
	}
	_, err = db.deleteCollectionMeta(collectionName)

	return err
}
