package spritedb

import (
	"fmt"

	"github.com/google/uuid"
)

type Query struct {
	db             *DB
	collectionName string
	slct           []string
	filters        []Filter
	skip           int
	take           int
	document       Document
}

func (query *Query) Insert() (string, error) {
	tx, err := query.db.db.Begin(true)
	if err != nil {
		return "-1", err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte(query.collectionName))
	if bucket == nil {
		return "-1", fmt.Errorf("collection %s not found", query.collectionName)
	}

	id := uuid.New().String()
	encodedValues, err := query.db.Encoder.Encode(query.document)

	if err != nil {
		return "-1", err
	}
	if err := bucket.Put([]byte(id), encodedValues); err != nil {
		return "-1", err
	}

	return id, tx.Commit()
}

func (query *Query) Select() ([]Document, error) {
	results := make([]Document, 0)
	tx, err := query.db.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte(query.collectionName))
	if bucket == nil {
		return nil, fmt.Errorf("collection %s not found", query.collectionName)
	}

	skip := 0
	c := bucket.Cursor()

	for k, v := c.First(); k != nil; k, v = c.Next() {
		if decodedDoc, err := query.db.Decoder.Decode(v); err != nil {
			return nil, err
		} else {
			decodedDoc["id"] = string(k)

			// only and works
			if !query.applyFilters(decodedDoc) {
				continue
			}
			if skip < query.skip {
				skip++
				continue
			}

			// select fields
			decodedDoc = query.SelectFields(decodedDoc)
			results = append(results, decodedDoc)
		}

		if len(results) == query.take {
			break
		}
	}

	return results, tx.Commit()
}

func (query *Query) UpdateDocument() (string, error) {
	tx, err := query.db.db.Begin(true)
	if err != nil {
		return "-1", nil
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte(query.collectionName))
	if bucket == nil {
		return "-1", nil
	}

	encodedValues, err := query.db.Encoder.Encode(query.document)

	if err != nil {
		return "-1", nil
	}
	if err := bucket.Put([]byte(query.document["id"]), encodedValues); err != nil {
		return "-1", nil
	}

	return query.document["id"], tx.Commit()
}

func (query *Query) DeleteDocumentById() (string, error) {
	var documentId = query.document["id"]
	tx, err := query.db.db.Begin(true)
	if err != nil {
		return "-1", nil
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte(query.collectionName))
	if bucket == nil {
		return "-1", nil
	}

	if err := bucket.Delete([]byte(documentId)); err != nil {
		return "-1", nil
	}

	return documentId, tx.Commit()
}

func (query *Query) DeleteDocumentsByDocument() (string, error) {

	tx, err := query.db.db.Begin(true)
	if err != nil {
		return "-1", nil
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte(query.collectionName))
	if bucket == nil {
		return "-1", nil
	}

	c := bucket.Cursor()
	for k, v := c.First(); k != nil; k, v = c.Next() {
		if decodedDoc, err := query.db.Decoder.Decode(v); err != nil {
			return "-1", err
		} else {
			decodedDoc["id"] = string(k)

			// only and works
			if !query.applyFilters(decodedDoc) {
				continue
			}

			if err := bucket.Delete([]byte(decodedDoc["id"])); err != nil {
				return "-1", err
			}
		}
	}

	return "1", tx.Commit()
}
