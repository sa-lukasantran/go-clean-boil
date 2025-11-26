package adapters

import (
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *Database

func GetDatabase() *Database {
	return DB
}

func StorageInit() {
	DB = &Database{InitDatabase()}
}
