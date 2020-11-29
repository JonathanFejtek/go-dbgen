package dbgen

import (
	"fmt"
	"testing"
)

func Test_NewInsert(t *testing.T) {

	type args struct {
		insertQuery InsertQuery
	}

	user := struct {
		ID        string `db:"id"`
		Name      string `db:"name"`
		Email     string `db:"email"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}{}
	tests := []struct {
		name            string
		args            args
		wantQueryString string
	}{
		{
			name: "basic query",
			args: args{
				insertQuery: NewInsert("users", user),
			},

			wantQueryString: fmt.Sprintf(
				templInsert,
				"users",
				"id, name, email, created_at, updated_at",
				":id, :name, :email, :created_at, :updated_at",
				"users.id, users.name, users.email, users.created_at, users.updated_at"),
		},
		{
			name: "omit insert values",
			args: args{
				insertQuery: NewInsert("users", user).OmitValues("created_at", "updated_at"),
			},

			wantQueryString: fmt.Sprintf(
				templInsert,
				"users",
				"id, name, email",
				":id, :name, :email",
				"users.id, users.name, users.email, users.created_at, users.updated_at"),
		},

		{
			name: "omit return values",
			args: args{
				insertQuery: NewInsert("users", user).OmitReturns("created_at", "updated_at"),
			},

			wantQueryString: fmt.Sprintf(
				templInsert,
				"users",
				"id, name, email, created_at, updated_at",
				":id, :name, :email, :created_at, :updated_at",
				"users.id, users.name, users.email"),
		},

		{
			name: "omit return values and insert values",
			args: args{
				insertQuery: NewInsert("users", user).
					OmitReturns("created_at", "updated_at").
					OmitValues("created_at", "updated_at"),
			},

			wantQueryString: fmt.Sprintf(
				templInsert,
				"users",
				"id, name, email",
				":id, :name, :email",
				"users.id, users.name, users.email"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			qString := tt.args.insertQuery.String()
			if qString != tt.wantQueryString {
				t.Errorf("Insert string = %+v ||  \n want %+v", qString, tt.wantQueryString)
			}
		})
	}
}

func Test_NewUpdate(t *testing.T) {

	type args struct {
		updateQuery UpdateQuery
	}

	user := struct {
		ID        string `db:"id"`
		Name      string `db:"name"`
		Email     string `db:"email"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}{}
	tests := []struct {
		name            string
		args            args
		wantQueryString string
	}{
		{
			name: "basic query",
			args: args{
				updateQuery: NewUpdate("users", user),
			},

			wantQueryString: fmt.Sprintf(
				templUpdate,
				"users",
				"id=:id, name=:name, email=:email, created_at=:created_at, updated_at=:updated_at",
				"id=:id",
				"users.id, users.name, users.email, users.created_at, users.updated_at"),
		},
		{
			name: "basic query custom where",
			args: args{
				updateQuery: NewUpdate("users", user).Where("email=:email AND id=:id"),
			},

			wantQueryString: fmt.Sprintf(
				templUpdate,
				"users",
				"id=:id, name=:name, email=:email, created_at=:created_at, updated_at=:updated_at",
				"email=:email AND id=:id",
				"users.id, users.name, users.email, users.created_at, users.updated_at"),
		},
		{
			name: "omit values",
			args: args{
				updateQuery: NewUpdate("users", user).OmitValues("id", "created_at", "updated_at"),
			},

			wantQueryString: fmt.Sprintf(
				templUpdate,
				"users",
				"name=:name, email=:email",
				"id=:id",
				"users.id, users.name, users.email, users.created_at, users.updated_at"),
		},
		{
			name: "omit returns",
			args: args{
				updateQuery: NewUpdate("users", user).OmitReturns("created_at", "updated_at"),
			},

			wantQueryString: fmt.Sprintf(
				templUpdate,
				"users",
				"id=:id, name=:name, email=:email, created_at=:created_at, updated_at=:updated_at",
				"id=:id",
				"users.id, users.name, users.email"),
		},
		{
			name: "omit returns and values",
			args: args{
				updateQuery: NewUpdate("users", user).OmitValues("id").
					OmitReturns("created_at", "updated_at"),
			},

			wantQueryString: fmt.Sprintf(
				templUpdate,
				"users",
				"name=:name, email=:email, created_at=:created_at, updated_at=:updated_at",
				"id=:id",
				"users.id, users.name, users.email"),
		},
		{
			name: "omit returns and values with custom where",
			args: args{
				updateQuery: NewUpdate("users", user).
					OmitValues("id").
					OmitReturns("created_at", "updated_at").
					Where("email=:x.f@xx.com"),
			},

			wantQueryString: fmt.Sprintf(
				templUpdate,
				"users",
				"name=:name, email=:email, created_at=:created_at, updated_at=:updated_at",
				"email=:x.f@xx.com",
				"users.id, users.name, users.email"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			qString := tt.args.updateQuery.String()
			if qString != tt.wantQueryString {
				t.Errorf("Insert string = %+v ||  \n want %+v", qString, tt.wantQueryString)
			}
		})
	}
}

func Test_NewGet(t *testing.T) {

	type args struct {
		getQuery GetQuery
	}

	user := struct {
		ID        string `db:"id"`
		Name      string `db:"name"`
		Email     string `db:"email"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}{}
	tests := []struct {
		name            string
		args            args
		wantQueryString string
	}{
		{
			name: "basic query",
			args: args{
				getQuery: NewGet("users", user),
			},

			wantQueryString: fmt.Sprintf(
				templSelect,
				"users.id, users.name, users.email, users.created_at, users.updated_at",
				"users",
				"id=:id",
			),
		},
		{
			name: "omit returns",
			args: args{
				getQuery: NewGet("users", user).OmitReturns("name", "created_at"),
			},

			wantQueryString: fmt.Sprintf(
				templSelect,
				"users.id, users.email, users.updated_at",
				"users",
				"id=:id",
			),
		},
		{
			name: "omit returns custom where",
			args: args{
				getQuery: NewGet("users", user).
					OmitReturns("name", "created_at").
					Where("email=:email"),
			},

			wantQueryString: fmt.Sprintf(
				templSelect,
				"users.id, users.email, users.updated_at",
				"users",
				"email=:email",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			qString := tt.args.getQuery.String()
			if qString != tt.wantQueryString {
				t.Errorf("Insert string = %+v ||  \n want %+v", qString, tt.wantQueryString)
			}
		})
	}
}

func Test_NewDelete(t *testing.T) {

	type args struct {
		deleteQuery DeleteQuery
	}

	user := struct {
		ID        string `db:"id"`
		Name      string `db:"name"`
		Email     string `db:"email"`
		CreatedAt string `db:"created_at"`
		UpdatedAt string `db:"updated_at"`
	}{}
	tests := []struct {
		name            string
		args            args
		wantQueryString string
	}{
		{
			name: "basic query",
			args: args{
				deleteQuery: NewDelete("users", user),
			},

			wantQueryString: fmt.Sprintf(
				templDelete,
				"users",
				"id=:id",
			),
		},

		{
			name: "custom where",
			args: args{
				deleteQuery: NewDelete("users", user).Where("id=1"),
			},

			wantQueryString: fmt.Sprintf(
				templDelete,
				"users",
				"id=1",
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			qString := tt.args.deleteQuery.String()
			if qString != tt.wantQueryString {
				t.Errorf("Insert string = %+v ||  \n want %+v", qString, tt.wantQueryString)
			}
		})
	}
}
