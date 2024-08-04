package spritedb

type QueryBuilder struct {
	db             *DB
	collectionName string
	slct           []string
	filters        []Filter
	skip           int
	take           int
	document       Document
}

func (db *DB) NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		db: db,
	}
}

func (qb *QueryBuilder) Collection(collectionName string) *QueryBuilder {
	qb.collectionName = collectionName
	return qb
}

func (qb *QueryBuilder) Insert(document Document) *QueryBuilder {
	qb.document = document
	return qb
}

func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
	qb.slct = fields
	return qb
}

func (qb *QueryBuilder) Where(filter Filter) *QueryBuilder {
	qb.filters = append(qb.filters, filter)
	return qb
}

func (qb *QueryBuilder) Skip(count int) *QueryBuilder {
	qb.skip = count
	return qb
}

func (qb *QueryBuilder) Take(count int) *QueryBuilder {
	qb.take = count
	return qb
}

func (qb *QueryBuilder) DeleteDocumentById(documentId string) *QueryBuilder {
	qb.document["id"] = documentId
	return qb
}

func (qb *QueryBuilder) Build() *Query {
	query := &Query{
		db:             qb.db,
		collectionName: qb.collectionName,
		slct:           qb.slct,
		filters:        qb.filters,
		skip:           qb.skip,
		take:           qb.take,
		document:       qb.document,
	}
	return query
}
