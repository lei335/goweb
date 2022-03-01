package godb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Album 表结构
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func connectDB() error {
	// 捕获连接属性
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	// 获得数据库句柄
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return pingErr
	}
	fmt.Println("Connected!")

	return nil
}

func doDB() {
	// 查询
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// 单行查询
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	// 插入数据
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	// 更新数据
	rows, err := updateAlbum(Album{
		6, "The star lost", "Monica", 66.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("rows of updated album: %v\n", rows)

	// 删除数据
	rows, err = deleteAlbum(albID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("rows of deleted album: %v\n", rows)
}

// 查询
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("=albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

// 单行查询
func albumByID(id int64) (Album, error) {
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumById %d: %v", id, err)
	}
	return alb, nil
}

// 插入数据
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// 更新数据
func updateAlbum(alb Album) (int64, error) {
	result, err := db.Exec("UPDATE album SET title=?, artist=?, price=? WHERE id=?", alb.Title, alb.Artist, alb.Price, alb.ID)
	if err != nil {
		return 0, fmt.Errorf("updateAlbum: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateAlbum: %v", err)
	}
	return rows, nil
}

// 删除数据
func deleteAlbum(id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM album WHERE id=?", id)
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %v", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %v", err)
	}
	return rows, nil
}
