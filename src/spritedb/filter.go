package spritedb

type Filter struct {
	field string
	Op    string
	Value string
}

func (f *Filter) applyOperation(document Document) bool {
	if f.Op == "eq" {
		return f.Value == document[f.field]
	} else if f.Op == "ne" {
		return f.Value != document[f.field]
	} else if f.Op == "gt" {
		return f.Value > document[f.field]
	} else if f.Op == "lt" {
		return f.Value < document[f.field]
	} else if f.Op == "gte" {
		return f.Value >= document[f.field]
	} else if f.Op == "lte" {
		return f.Value <= document[f.field]
	} else if f.Op == "like" {
		return f.Value == document[f.field]
	}
	return false
}

func (query Query) applyFilters(document Document) bool {
	for _, filter := range query.filters {
		if !filter.applyOperation(document) {
			return false
		}
	}
	return true
}
