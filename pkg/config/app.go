package config

import(
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)


var (
	db *gorm.DB
)




func Connect(){
	godotenv.Load(".env")
	MysqlConnect := os.Getenv("MYSQLCONNECT")
	MysqlConnect = fmt.Sprint(MysqlConnect + "?charset=utf8&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", MysqlConnect)

	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}