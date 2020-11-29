package dbgen

import (
	"fmt"
)

// SelectQuerier interface required to build a select rows db function
type SelectQuerier interface {
	Select(query string, dest interface{}, args ...interface{}) error
}

// SelectQuerier interface required to build a select row db function
type SelectOneQuerier interface {
	SelectOne(query string, i interface{}, args ...interface{}) error
}

// MakeUpdateQueryArgs arguments required to make an update query
type MakeGetQueryArgs struct {
	TableName    string
	WhereClause  string
	ReturnFields Columns
}

// GetQuery represents a get query
type GetQuery struct {
	query
	makeQuery func(args MakeGetQueryArgs) string
}

// OmitReturns omit return/select fields from the get query
func (q GetQuery) OmitReturns(fields ...string) GetQuery {
	nq := q
	nq.query = nq.query.omitReturns(fields...)
	return nq
}

// Where set the where clause of the get query
func (q GetQuery) Where(whereString string) GetQuery {
	nq := q
	nq.query = nq.query.where(whereString)
	return nq
}

// String generate the get query as a string query
func (q GetQuery) String() string {

	return q.makeQuery(MakeGetQueryArgs{
		TableName:    q.tableName,
		WhereClause:  q.whereClause,
		ReturnFields: q.returnFields,
	})

}

// FnSelect generate the get query as a function to select multiple rows from a DB
func (q GetQuery) FnSelect() func(tx SelectQuerier, i interface{}, args ...interface{}) error {
	qs := q.String()
	return func(tx SelectQuerier, i interface{}, args ...interface{}) error {
		return tx.Select(qs, i, args...)
	}
}

// FnSelectOne generate the get query as a function to select a single row from a DB
func (q GetQuery) FnSelectOne() func(tx SelectOneQuerier, i interface{}, args ...interface{}) error {
	qs := q.String()
	return func(tx SelectOneQuerier, i interface{}, args ...interface{}) error {
		return tx.SelectOne(qs, i, args...)
	}
}

// GetQueryOptions optional arguments to create a new get query
type GetQueryOptions struct {
	MakeQuery func(args MakeGetQueryArgs) string
}

// NewGet generate a new get query
func NewGet(tableName string, i interface{}, opts ...GetQueryOptions) GetQuery {
	tags := getTagsByName("db", i)

	var options GetQueryOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	q := GetQuery{
		query: query{
			tableName: tableName,
			returnFields: Columns{
				TableName: tableName,
				Fields:    tags,
			},
			whereClause: DefaultIdentityString,
		},

		makeQuery: func(args MakeGetQueryArgs) string {
			return fmt.Sprintf(
				templSelect,
				args.ReturnFields.AsSelects().Joined(),
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
