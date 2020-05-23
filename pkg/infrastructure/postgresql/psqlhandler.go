package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"otter/pkg/interfaces/gateway"
)

type PsqlHandler struct {
	Conn *sql.DB
}

func (handler *PsqlHandler) Execute(statement string, args ...interface{}) (gateway.Result, error) {
	psqlResult := PsqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return psqlResult, err
	}
	psqlResult.Result = result
	return psqlResult, err
}

func (handler *PsqlHandler) Query(statement string, args ...interface{}) (gateway.Row, error) {
	psqlRow := new(PsqlRow)
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return psqlRow, err
	}
	psqlRow.Rows = rows
	return psqlRow, err
}

type PsqlResult struct {
	Result sql.Result
}

func (result PsqlResult) LastInsertId() (int64, error) {
	return result.Result.LastInsertId()
}

func (result PsqlResult) RowsAffected() (int64, error) {
	return result.Result.RowsAffected()
}

type PsqlRow struct {
	Rows *sql.Rows
}

func (row PsqlRow) Scan(dest ...interface{}) error {
	return row.Rows.Scan(dest...)
}

func (row PsqlRow) Next() bool {
	return row.Rows.Next()
}

func (row PsqlRow) Close() error {
	return row.Rows.Close()
}

func NewPsqlHandler(dbfileName string) gateway.SqlHandler {
	conn, err := sql.Open("postgres", dbfileName)
	if err != nil {
		panic(err.Error)
	}
	psqlHandler := new(PsqlHandler)
	psqlHandler.Conn = conn
	return psqlHandler
}
