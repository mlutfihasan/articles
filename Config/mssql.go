package Config

import (
	"database/sql"
	"log"

	mssql "github.com/denisenkom/go-mssqldb"
)

func Connect() (db *sql.DB) {
	connector, err := mssql.NewConnector("sqlserver://app_misplus:VNEU-MF-2022@103.103.192.24:21033/mfdb")
	if err != nil {
		log.Println(err)
		db = nil
	}
	connector.SessionInitSQL = "SET ANSI_NULLS ON"
	db = sql.OpenDB(connector)

	return db
}
