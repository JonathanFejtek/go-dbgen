package dbgen

import (
	"fmt"
)

// DeleteQuerier interface that needs to be satisfied to construct a delete db function
type DeleteQuerier interface {
	Delete(q string, args ...interface{}) (int64, error)
}

// MakeDeleteQueryArgs arguments required to make a delete query
type MakeDeleteQueryArgs struct {
	TableName   string
	WhereClause string
}

// DeleteQuery represents a delete query
type DeleteQuery struct {
	query
	makeQuery func(args MakeDeleteQueryArgs) string
}

// String the delete query as a string query
func (q DeleteQuery) String() string {
	return q.makeQuery(MakeDeleteQueryArgs{
		TableName:   q.tableName,
		WhereClause: q.whereClause,
	})
}

// Where set the where clause of the delete query
func (q DeleteQuery) Where(whereString string) DeleteQuery {
	nq := q
	nq.query = nq.query.where(whereString)
	return nq
}

// Fn generate a db delete function
func (q DeleteQuery) Fn() func(tx DeleteQuerier, args ...interface{}) (int64, error) {
	qs := q.String()
	return func(tx DeleteQuerier, args ...interface{}) (int64, error) {
		return tx.Delete(qs, args...)
	}
}

// DeleteQueryOptions optional arguments to create a new delete query
type DeleteQueryOptions struct {
	MakeQuery func(args MakeDeleteQueryArgs) string
}

// NewDelete construct a new delete query
func NewDelete(tableName string, i interface{}, opts ...DeleteQueryOptions) DeleteQuery {

	var options DeleteQueryOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	q := DeleteQuery{
		query: query{
			tableName:   tableName,
			whereClause: DefaultIdentityString,
		},
		makeQuery: func(args MakeDeleteQueryArgs) string {
			return fmt.Sprintf(
				templDelete,
				args.TableName,
				args.WhereClause,
			)
		},
	}

	if options.MakeQuery != nil {
		q.makeQuery = options.MakeQuery
	}

	return q
}
