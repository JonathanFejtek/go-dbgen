package dbgen

import (
	"fmt"
)

// UpdateQuerier interface required to build update db functions
type UpdateQuerier interface {
	Update(q string, val interface{}) error
}

// MakeUpdateQueryArgs arguments required to make an update query
type MakeUpdateQueryArgs struct {
	TableName    string
	Values       Columns
	WhereClause  string
	ReturnFields Columns
}

// UpdateQuery represents an update query
type UpdateQuery struct {
	query
	makeQuery func(args MakeUpdateQueryArgs) string
}

// OmitValues omit values to update from the query
func (q UpdateQuery) OmitValues(fields ...string) UpdateQuery {
	nq := q
	nq.query = nq.query.omitValues(fields...)
	return nq
}

// OmitValues omit fields to return from the query
func (q UpdateQuery) OmitReturns(fields ...string) UpdateQuery {
	nq := q
	nq.query = nq.query.omitReturns(fields...)
	return nq
}

// Where set the where clause of the query
func (q UpdateQuery) Where(whereString string) UpdateQuery {
	nq := q
	nq.query = nq.query.where(whereString)
	return nq
}

// String generate the query as a string
func (q UpdateQuery) String() string {

	return q.makeQuery(
		MakeUpdateQueryArgs{
			TableName:    q.tableName,
			Values:       q.valueFields,
			WhereClause:  q.whereClause,
			ReturnFields: q.returnFields,
		},
	)
}

// Fn generate the query as a function
func (q UpdateQuery) Fn() func(tx UpdateQuerier, i interface{}) error {
	qs := q.String()
	return func(tx UpdateQuerier, i interface{}) error {
		return tx.Update(qs, i)
	}
}

// UpdateQueryOptions optional arguments to create a new update query
type UpdateQueryOptions struct {
	MakeQuery func(args MakeUpdateQueryArgs) string
}

// NewUpdate construct a new update query
func NewUpdate(
	tableName string,
	i interface{},
	opts ...UpdateQueryOptions,
) UpdateQuery {
	tags := getTagsByName("db", i)

	var options UpdateQueryOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	q := UpdateQuery{
		query: query{
			tableName: tableName,
			valueFields: Columns{
				TableName: tableName,
				Fields:    tags,
			},
			returnFields: Columns{
				TableName: tableName,
				Fields:    tags,
			},
			whereClause: DefaultIdentityString,
		},
		makeQuery: func(
			args MakeUpdateQueryArgs,
		) string {
			return fmt.Sprintf(
				templUpdate,
				args.TableName,
				args.Values.AsAssignments().Joined(),
				args.WhereClause,
				args.ReturnFields.AsSelects().Joined(),
			)
		},
	}

	if options.MakeQuery != nil {
		q.makeQuery = options.MakeQuery
	}

	return q
}
