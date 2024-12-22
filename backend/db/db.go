package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB; 

func InitDB() {
    dsn := os.Getenv("DATABASE_URI");
    var err error;
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{});

    if err != nil {
        log.Fatal(err)
        return
    }
}

func GetDB() *gorm.DB {
    return db;
}
