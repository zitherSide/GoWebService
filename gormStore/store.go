package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    //XxIdはgormが外部キーとして扱ってくれる
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, &Comment{})
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "良い投稿だね！", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(comment) //DB.Modelで街頭レコードのモデルにアクセスできる

	var readPost Post
	Db.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
