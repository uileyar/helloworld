package main

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type Invoice struct {
	Id       int64
	Created  int64
	Updated  int64
	Memo     string
	PersonId int64
}

type Person struct {
	Id      int64
	Created int64
	Updated int64
	FName   string
	LName   string
}

// Example of using tags to alias fields to column names
// The 'db' value is the column name
//
// A hyphen will cause gorp to skip this field, similar to the
// Go json package.
//
// This is equivalent to using the ColMap methods:
//
//   table := dbmap.AddTableWithName(Product{}, "product")
//   table.ColMap("Id").Rename("product_id")
//   table.ColMap("Price").Rename("unit_price")
//   table.ColMap("IgnoreMe").SetTransient(true)
//
// You can optionally declare the field to be a primary key and/or autoincrement
//
type Product struct {
	Id       int64  `db:"product_id, primarykey, autoincrement"`
	Price    int64  `db:"unit_price"`
	IgnoreMe string `db:"-"`
}

func Mapping_structs_to_tables() {

	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, _ := sql.Open("mymysql", "tcp:localhost:3306*mydb/myuser/mypassword")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	// register the structs you wish to use with gorp
	// you can also use the shorter dbmap.AddTable() if you
	// don't want to override the table name
	//
	// SetKeys(true) means we have a auto increment primary key, which
	// will get automatically bound to your struct post-insert
	//
	dbmap.AddTableWithName(Invoice{}, "invoice_test").SetKeys(true, "Id")
	dbmap.AddTableWithName(Person{}, "person_test").SetKeys(true, "Id")
	dbmap.AddTableWithName(Product{}, "product_test").SetKeys(true, "Id")

	dbmap.CreateTablesIfNotExists()
}
