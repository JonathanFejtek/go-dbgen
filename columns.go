package dbgen

import (
	"fmt"
	"strings"
)

// Columns represents the columns of a db table
type Columns struct {
	TableName string
	Fields    []string
}

// Omit omit columns from the set
func (cc Columns) Omit(fields ...string) Columns {
	newFields := filterTags(cc.Fields, fields)
	return Columns{
		TableName: cc.TableName,
		Fields:    newFields,
	}
}

// Add add columns to the set
func (cc Columns) Add(fields ...string) Columns {
	var added []string
	added = append(added, cc.Fields...)
	added = append(added, fields...)

	return Columns{
		TableName: cc.TableName,
		Fields:    added,
	}
}

// Set set the columns
func (cc Columns) Set(fields ...string) Columns {
	var added []string
	added = append(added, fields...)

	return Columns{
		TableName: cc.TableName,
		Fields:    added,
	}
}

// Joined get the column set as a joined string
func (cc Columns) Joined() string {
	return strings.Join(cc.Fields, ", ")
}

// AsSelects get the columns formatted to be selected from their parent table
// suitable for SELECT and RETURNING statements
func (cc Columns) AsSelects() Columns {
	var params []string
	for _, c := range cc.Fields {
		params = append(
			params,
			fmt.Sprintf(
				"%s.%s", cc.TableName, c,
			),
		)
	}

	return Columns{
		TableName: cc.TableName,
		Fields:    params,
	}
}

// AsParams get the columns formatted as query template parameters
func (cc Columns) AsParams() Columns {
	var params []string
	for _, c := range cc.Fields {
		params = append(
			params,
			fmt.Sprintf(
				":%s", c,
			),
		)
	}

	return Columns{
		TableName: cc.TableName,
		Fields:    params,
	}
}

// AsAssignments get the columns formatted as query template assignments
func (cc Columns) AsAssignments() Columns {
	var params []string
	for _, c := range cc.Fields {
		params = append(
			params,
			fmt.Sprintf(
				"%s=:%s", c, c,
			),
		)
	}

	return Columns{
		TableName: cc.TableName,
		Fields:    params,
	}
}

func NewFieldBuilder(tableName string, i interface{}) Columns {
	tags := getTagsByName("db", i)
	return Columns{
		TableName: tableName,
		Fields:    tags,
	}
}
