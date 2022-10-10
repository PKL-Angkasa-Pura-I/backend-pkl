package database

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/config"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)
	DB, err := gorm.Open(mysql.Open(conectionString), &gorm.Config{})
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	DB.AutoMigrate(&model.Division{})
	return DB
}
