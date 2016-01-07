package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-gorp/gorp"
)

/*
 [account_id] VARCHAR(36) NOT NULL,
  [server_id] INTEGER DEFAULT (0),
  [user_id] integer DEFAULT (0),
  [name] nvarchar(10) NOT NULL,
  [kind] nvarchar(10) NOT NULL,
  [amount] float NOT NULL,
  [description] NVARCHAR(200),
  [sort] INTEGER DEFAULT (0),
  [version] TIMESTAMP DEFAULT (CURRENT_TIMESTAMP),
  CONSTRAINT [sqlite_autoindex_jzb_accounts_1] PRIMARY KEY ([account_id])
*/
type Account struct {
	Account_id  string `db:", size:36, primarykey"`
	Server_id   int
	User_id     int
	Name        string `db:", size:10"`
	Kind        string `db:", size:10"`
	Amount      float32
	Description string `db:", size:200"`
	Sort        int
	Version     time.Time `db:", default:CURRENT_TIMESTAMP"`
}

func (u *Account) PreInsert(s gorp.SqlExecutor) error {
	//u.Account_id = CreateGUID()
	//u.Version = time.Now()

	var val int
	if err := s.SelectOne(&val, "select max(sort) from jzb_accounts"); err == nil {
		u.Sort = val + 1
	}
	fmt.Println(u)
	return nil
}

func (u *Account) PreUpdate(s gorp.SqlExecutor) error {
	//u.Version = time.Now()
	return nil
}

func InsertTest() {
	db, err := sql.Open("sqlite3", "webapp.db")
	if err != nil {
		fmt.Println(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	defer dbmap.Db.Close()

	dbmap.AddTableWithName(Account{}, "jzb_accounts")
	//err = dbmap.CreateTablesIfNotExists()

	account := &Account{
		Account_id: "11211fwr",
		Name:       "1111",
		Kind:       "1111",
	}
	err = dbmap.Insert(account)
	if err != nil {
		fmt.Println(err)
	}
}
