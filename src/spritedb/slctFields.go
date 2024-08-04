package spritedb

func (query *Query) SelectFields(document Document) Document {
	if len(query.slct) == 0 {
		return document
	}

	newDoc := make(Document, 0)
	for _, field := range query.slct {
		if val, ok := document[field]; ok {
			newDoc[field] = val
		} else {
			newDoc[field] = ""
		}

	}
	newDoc["id"] = document["id"]
	return newDoc
}
