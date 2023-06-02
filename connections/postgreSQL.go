package connections

import (
	"database/sql"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*sql.DB, error) {
	dns := "host=localhost user=root password=postgresql dbname=book-store port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	psqlDB, err := conn.DB()
	if err != nil {
		log.Println(err)
	}
	return psqlDB, nil
}
