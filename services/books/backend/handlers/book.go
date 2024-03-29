package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Properties struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

var LoadedProperties Properties

func LoadProperties() {
	err := godotenv.Load("./properties.env")
	if err != nil {
		log.Fatalf("error occured with the message: %s", err)
	}

	LoadedProperties.Host = os.Getenv("HOST")
	LoadedProperties.Dbname = os.Getenv("DB_DBNAME")
	LoadedProperties.Port = os.Getenv("DB_PORT")
	LoadedProperties.User = os.Getenv("DB_USER")
	LoadedProperties.Password = os.Getenv("DB_PASSWORD")
}

func ConnectToPSQL() (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s sslmode=disable",
		LoadedProperties.Host, LoadedProperties.Port, LoadedProperties.User, LoadedProperties.Password, LoadedProperties.Dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()

	return db, err
}

type Book struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	Author     string `json:"author"`
	ISBN       string `json:"isbn"`
	Edition    int    `json:"edition"`
	Year       int    `json:"year"`
	ShelfName  string `json:"shelf_name"`
	ShelfLevel int    `json:"shelf_level"`
}

type Books []*Book

type BookSlice struct {
	Books []*Book `json:"books"`
}
