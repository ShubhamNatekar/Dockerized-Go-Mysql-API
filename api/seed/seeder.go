package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ShubhamNatekar/Dockerized-Go-Mysql-API/api/models"
)

var users = []models.User{
	models.User{
		Name: "rakesh singh",
		Age: 23,
		Department:    "CS",
		Subject: "DOS",
	},
	models.User{
		Name: "bhavesh nadurdikar",
		Age: 21,
		Department:    "CS",
		Subject: "DS",
	},
	models.User{
		Name: "payal jain",
		Age: 22,
		Department:    "CS",
		Subject: "PDA",
	},
	models.User{
		Name: "esha varma",
		Age: 22,
		Department:    "CS",
		Subject: "DOS",
	},

}


func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Product{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}

