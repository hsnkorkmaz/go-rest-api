package db

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-rest-api/config"
)

var SQL_DB *sql.DB

func SQL_Query(query string) *sql.Rows {
	err := openSqlConnection()
	if err != nil {
		fmt.Println("Error:", err)
	}
	rows, err := SQL_DB.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return rows
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
