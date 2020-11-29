package dbgen

type query struct {
	tableName    string
	valueFields  Columns
	returnFields Columns
	whereClause  string
}

func (q query) omitValues(fields ...string) query {
	q2 := q
	q2.valueFields = q.valueFields.Omit(fields...)

	return q2
}

func (q query) omitReturns(fields ...string) query {
	q2 := q
	q2.returnFields = q2.returnFields.Omit(fields...)
	return q2
}

func (q query) where(whereString string) query {
	q2 := q
	q2.whereClause = whereString
	return q2
}
