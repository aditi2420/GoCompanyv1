package models

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB
type dbHandle struct{
	*gorm.DB
}

var db *dbHandle

func SetupDatabase() {
	
	 	// dsn :=  fmt.Sprintf(
 		// "host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable ", configs.DatabaseHost,configs.DatabaseUser,
 		// configs.DatabasePassword, configs.DatabaseName)

    dsn := fmt.Sprintf(
        "host=company_postgres user=%s password=%s dbname=%s port=5432 sslmode=disable",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    dbHandle1, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(1)
    }

    log.Println("connected")
    dbHandle1.Logger = logger.Default.LogMode(logger.Info)
	db = &dbHandle{ dbHandle1}
    log.Println("running migrations")
 	db.AutoMigrate(&Company{})
}

func  GetDbHandle() *dbHandle {
	return db
}