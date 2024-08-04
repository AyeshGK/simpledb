package spritedb

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CollectionMeta struct {
	ID        string
	Name      string
	CreatedAt string
}

func (db *DB) CreateCollectionMeta(collectionName string) (string, error) {
	tx, err := db.db.Begin(true)
	if err != nil {
		return "-1", err
	}
	defer tx.Rollback()
	// get bucket without opening transaction because transaction is already opened
	bucket := tx.Bucket([]byte("collection-meta"))
	if bucket == nil {
		return "-1", fmt.Errorf("ERROR: Database Error Meta Collection not found")
	}

	id := uuid.New().String()
	collectionMetaData := CollectionMeta{
		ID:        id,
		Name:      collectionName,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	encodedValues, err := json.Marshal(collectionMetaData)
	if err != nil {
		return "-1", err
	}

	if err := bucket.Put([]byte(id), encodedValues); err != nil {
		return "-1", err
	}

	return id, tx.Commit()
}

func (db *DB) getAllCollectionMeta() ([]Document, error) {
	query := db.NewQueryBuilder().Collection("collection-meta").Build()
	return query.Select()
}

func (db *DB) deleteCollectionMeta(collectionName string) (string, error) {
	query := db.NewQueryBuilder().Collection("collection-meta").DeleteDocumentById(collectionName).Build()
	return query.DeleteDocumentById()
}
