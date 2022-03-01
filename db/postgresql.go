package godb

import (
	"database/sql"
	"fmt"
)

type UserInfo struct {
	UID        int64
	UserName   string
	Department string
	Created    string
}

// ConnectPQ connect postgresql
func ConnectPQ() (*sql.DB, error) {

	dbDsn := "user=postgres password=123456 dbname=recordings sslmode=disable"
	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	fmt.Println("Connected!")

	return db, nil
}

func InsertData(user *UserInfo) error {
	db, err := ConnectPQ()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
	if err != nil {
		return err
	}

	res, err := stmt.Exec("alex", "产品部门", "2019-07-01")
	if err != nil {
		return err
	}

	var lastInsertId int
	err = db.QueryRow("", )
}
