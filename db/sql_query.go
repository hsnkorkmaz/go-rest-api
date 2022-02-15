package db

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-rest-api/config"
)

var SQL_DB *sql.DB

func SQL_Query(query string, args ...interface{}) (*sql.Rows, error) {
	err := openSqlConnection()
	if err != nil {
		return nil, err
	}
	rows, err := SQL_DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func openSqlConnection() error {
	var err error
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.SQL_USER, config.SQL_PASSWORD, config.SQL_SERVER, config.SQL_PORT, config.SQL_DATABASE)

	SQL_DB, err = sql.Open("mssql", connectionString)
	if err != nil {
		return err
	}
	return nil
}

func CloseSqlConnection() {
	SQL_DB.Close()
}
