package main

import (
	godb "go-web/db"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	_, err := godb.ConnectPQ()
	if err != nil {
		log.Fatal(err)
	}

	// 插入数据
	uinfo := &godb.UserInfo{
		UserName:   "Lily",
		Department: "UI部门",
		Created:    "2021-10-08",
	}
	lastInsertID, err := godb.InsertData(uinfo)
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	err = godb.QueryData()
	if err != nil {
		log.Fatal(err)
	}

	// 更新数据
	uinfo = &godb.UserInfo{
		UserName: "Bob",
		UID:      lastInsertID,
	}
	err = godb.UpdateData(uinfo)
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	err = godb.QueryData()
	if err != nil {
		log.Fatal(err)
	}
}
