package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // must import
	"github.com/volatiletech/sqlboiler/boil"
)

// Init initialize database connection
func Init() {
	db, err := sql.Open("mysql", "sampledb:sampledb@tcp([localhost]:3306)/sampledb?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	boil.SetDB(db)
	boil.DebugMode = true
}
