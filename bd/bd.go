package bd

import (
	"log"

	"github.com/jinzhu/gorm"
)

var Bd *gorm.DB

type Feed struct {
	ID       uint   `gorm:"primary_key"`
	URL      string `gorm:"type:nvarchar(1024);unique_index;not null"`
	Standard bool
}

func (Feed) TableName() string {
	return "Feeds"
}

type User struct {
	gorm.Model
	ID uint `gorm:"primary_key"`
}

func (User) TableName() string {
	return "Users"
}

type userFeed struct {
	UserID     uint   `gorm:"primary_key"`
	FeedID     uint   `gorm:"primary_key"`
	Desription string `gorm:"type:nvarchar(1024);unique_index;not null"`

	User User `gorm:"ForeignKey:UserID"`
	Feed Feed `gorm:"ForeignKey:FeedlID"`
}

func (userFeed) TableName() string {
	return "User_Feeds"
}

func init() {
	var err error
	Bd, err = gorm.Open("mysql", "GO_mysql_connector:L65gUIfd7i9JGHr4jhgH@(127.0.0.1:3306)/RSS_agregator_telegram_bot?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	Bd.LogMode(true)
}

func CreateFeed(url string, fl bool) {
	Bd.Create(&Feed{URL: url, Standard: fl})
}

func Select(id uint) []int {
	var idU []int
	Bd.Table("Users").Select("Users.id").Joins("JOIN User_Feeds JOIN Feeds ON Users.id = User_Feeds.`user` AND User_Feeds.feed = Feeds.id").Where("Feeds.id = ?", id).Pluck("Users.id", &idU)

	return idU
}

func MyPluck() []string {
	var urlF []string
	err := Bd.Table("Feeds").Pluck("Feeds.url", &urlF).Error
	if err != nil {
		log.Fatal(err)
	}
	return urlF
}