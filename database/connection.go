package database

import (
	"fmt"
	"log"
	"main/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func createDatabase(connection *gorm.DB, dbName string) error {
	var exists bool
	err := connection.Raw("SELECT COUNT(*) > 0 FROM information_schema.schemata WHERE schema_name = ?", dbName).Scan(&exists).Error
	if err != nil {
		return fmt.Errorf("error checking if database exists: %v", err)
	}

	if !exists {
		err := connection.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
		if err != nil {
			return fmt.Errorf("error creating database: %v", err)
		}
	}

	return nil
}

func Connect() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve database credentials from the environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true", dbUser, dbPassword, dbHost, dbPort)

	// Important to parseTime=true else will have error accessing posts
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to mysql")
	}

	err = createDatabase(connection, dbName)
	if err != nil {
		panic("error creating database")
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	connection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the db")
	}

	DB = connection

	// user
	connection.AutoMigrate(&models.User{})

	// threads
	connection.AutoMigrate(&models.Post{})

	// comments
	connection.AutoMigrate(&models.Comment{})
}
