package godb

import (
	"database/sql"
	"fmt"
)

var pgdb *sql.DB

type UserInfo struct {
	UID        int
	UserName   string
	Department string
	Created    string
}

// ConnectPQ connect postgresql
func ConnectPQ() (*sql.DB, error) {
	var err error

	dbDsn := "user=postgres password=123456 dbname=recordings sslmode=disable"
	pgdb, err = sql.Open("postgres", dbDsn)
	if err != nil {
		return nil, err
	}

	pingErr := pgdb.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	fmt.Println("Connected!")

	return pgdb, nil
}

func InsertData(user *UserInfo) (int, error) {

	stmt, err := pgdb.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
	if err != nil {
		return 0, err
	}

	_, err = stmt.Exec(user.UserName, user.Department, user.Created)
	if err != nil {
		return 0, err
	}

	_, err = stmt.Exec("alex", "产品部门", "2019-07-01")
	if err != nil {
		return 0, err
	}

	var lastInsertID int
	// 会执行插入操作并返回lastInsertID
	err = pgdb.QueryRow("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) returning uid;", "alex", "产品部门", "2019-07-01").Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}
	fmt.Println("最后插入id = ", lastInsertID)
	return lastInsertID, nil
}

func UpdateData(user *UserInfo) error {

	stmt, err := pgdb.Prepare("update userinfo set username=$1 where uid=$2")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(user.UserName, user.UID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("update affected nums: ", affect)
	return nil
}

func QueryData() error {

	rows, err := pgdb.Query("SELECT * FROM userinfo")
	if err != nil {
		return err
	}

	uinfo := &UserInfo{}
	for rows.Next() {
		err = rows.Scan(&uinfo.UID, &uinfo.UserName, &uinfo.Department, &uinfo.Created)
		if err != nil {
			return err
		}
		fmt.Printf("uid: %v, username: %v, department: %v, created: %v\n", uinfo.UID, uinfo.UserName, uinfo.Department, uinfo.Created)
	}
	return nil
}
