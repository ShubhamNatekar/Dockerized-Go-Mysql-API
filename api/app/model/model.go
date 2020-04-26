package model
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type Employee struct {
	gorm.Model
	ID    int    `json:"id"`
	Name   string `gorm:"unique" json:"name"`
	Age    int    `json:"age"`
	Dept   string `json:"dept"`
	Subject   string `json:"subject"`
}
// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}
