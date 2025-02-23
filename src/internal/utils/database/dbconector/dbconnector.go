package dbconector

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DataBase struct {
	DB *sql.DB
}

func NewDBConnector(db *sql.DB) *DataBase {
	return &DataBase{
		DB: db,
	}
}

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewConfig() *Config {
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}

func ConnectToDB(c *Config) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}

func MigrateDB(c *Config) error {

	connDsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)

	m, err := migrate.New(
		"file://internal/utils/database/migrations/",
		connDsn,
	)
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return fmt.Errorf("error applying migration: %v", err)
	}
	fmt.Println("Migration applied successfully!")
	return nil
}
