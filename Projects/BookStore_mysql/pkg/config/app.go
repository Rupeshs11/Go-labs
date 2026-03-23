package config

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	fmt.Println("Connecting to MySQL database...")
	d, err := gorm.Open("mysql","root:Technopath256@tcp(127.0.0.1:3306)/go_bookstore?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s")
	if err !=nil{
		panic(err)
	}
	fmt.Println("Successfully connected to database!")
	db = d
}

func GetDB() *gorm.DB{
	return db
}