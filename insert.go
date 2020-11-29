package dbgen

import "fmt"

// InsertQuerier interface required to build an insert db function
type InsertQuerier interface {
	Insert(q string, val interface{}) error
}

// MakeUpdateQueryArgs arguments required to make an update query
type MakeInsertQueryArgs struct {
	TableName    string
	Values       Columns
	ReturnFields Columns
}

// InsertQuery represents an insert query
type InsertQuery struct {
	query     query
	makeQuery func(args MakeInsertQueryArgs) string
}

// OmitValues omit value fields to insert from the query
func (q InsertQuery) OmitValues(fields ...string) InsertQuery {

	iq := q
	// iq := InsertQuery{
	// 	query: q.query.omitValues(fields...),
	// }

	iq.query = iq.query.omitValues(fields...)
	return iq
}

// OmitReturns omit return fields from the insert query
func (q InsertQuery) OmitReturns(fields ...string) InsertQuery {

	iq := q
	// iq := InsertQuery{
	// 	query: q.query.omitReturns(fields...),
	// }

	iq.query = iq.query.omitReturns(fields...)
	return iq
}

// String generate the query as a string
func (q InsertQuery) String() string {
	return q.makeQuery(
		MakeInsertQueryArgs{
			TableName:    q.query.tableName,
			Values:       q.query.valueFields,
			ReturnFields: q.query.returnFields,
		},
	)
}

// String generate the query as db function
func (q InsertQuery) Fn() func(tx InsertQuerier, i interface{}) error {
	qs := q.String()
	return func(tx InsertQuerier, i interface{}) error {
		return tx.Insert(qs, i)
	}
}

// InsertQueryOptions optional arguments to create a new insert query
type InsertQueryOptions struct {
	MakeQuery func(args MakeInsertQueryArgs) string
}

// NewInsert construct a new insert query
func NewInsert(
	tableName string,
	i interface{},
	opts ...InsertQueryOptions,
) InsertQuery {
	tags := getTagsByName("db", i)

	var options InsertQueryOptions
	if len(opts) > 0 {
		options = opts[0]
	}

	iq := InsertQuery{
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
		},
		makeQuery: func(
			args MakeInsertQueryArgs,
		) string {
			return fmt.Sprintf(
				templInsert,
				args.TableName,
				args.Values.Joined(),
				args.Values.AsParams().Joined(),
				args.ReturnFields.AsSelects().Joined(),
			)
		},
	}

	if options.MakeQuery != nil {
		iq.makeQuery = options.MakeQuery
	}

	return iq

}
