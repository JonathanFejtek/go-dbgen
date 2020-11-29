package dbgen

var (
	DefaultIdentityString = "id=:id"
)

const (
	templSelect = `SELECT %s FROM %s WHERE %s`

	templInsert = `INSERT INTO %s (
		%s
	) VALUES (
		%s
	)
	RETURNING %s`

	templUpdate = `UPDATE %s
	SET
		%s
	WHERE %s
	RETURNING %s`

	templDelete = `DELETE FROM %s WHERE %s`
)
