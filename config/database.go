package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

func buildDBConfig(host, port, user, name, password string) *DBConfig {
	dbConfig := DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		DBName:   name,
		Password: password,
	}
	return &dbConfig
}

func dbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func InitializeDatabase() (DB *gorm.DB) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dbConfig := buildDBConfig(dbHost, dbPort, dbUser, dbName, dbPassword)
	dbURL := dbURL(dbConfig)

	db, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Printf("Cannot connect to database. Host name: %s", dbHost)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", "mysql")
	}

	return db
}
