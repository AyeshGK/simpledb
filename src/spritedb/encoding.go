package spritedb

import (
	"encoding/json"
	"fmt"
)

type DataEncoder interface {
	Encode(Document) ([]byte, error)
}

type DataDecoder interface {
	Decode([]byte) (Document, error)
}

type JSONEncoder struct{}

func (e *JSONEncoder) Encode(document Document) ([]byte, error) {
	return json.Marshal(document)
}

type JSONDecoder struct{}

func (d *JSONDecoder) Decode(data []byte) (Document, error) {
	var document Document
	err := json.Unmarshal(data, &document)
	if err != nil {
		return nil, fmt.Errorf("error decoding document: %w", err)
	}
	return document, nil
}
