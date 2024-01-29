package mysql

import (
	"fmt"
	"github.com/ghiac/go-commons/log"
	"github.com/ghiac/social/internal/repositories"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Option struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}

func New(option Option) *repositories.Database {
	url := fmt.Sprintf("%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&charset=utf8mb4", option.User, option.Pass, option.Host, option.Port, option.Db)
	connection, err := gorm.Open("mysql", url)

	connection = connection.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8 auto_increment=1")

	//defer connection.Close()

	connection.AutoMigrate(&repositories.User{})
	connection.AutoMigrate(&repositories.Tweet{})
	connection.AutoMigrate(&repositories.Reaction{})
	connection.AutoMigrate(&repositories.Relation{})
	connection.AutoMigrate(&repositories.Session{})
	connection.AutoMigrate(&repositories.Event{})
	connection.AutoMigrate(&repositories.Hashtag{})

	if err != nil {
		log.Logger.Fatal("failed to connect database", err)
	}
	connection.LogMode(false)

	return &repositories.Database{
		Connection: connection,
	}
}
