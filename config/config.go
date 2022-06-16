package config

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ConnectDb() *sqlx.DB {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// load environment variables
	err := godotenv.Load(".env")

	// handle error, if any
	if err != nil {
		_ = fmt.Errorf("error loading .env file")
	}

	// format connection string from environment variables
	databaseURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
	)

	// create a database object which can be used to communicate with database
	db, err := sqlx.ConnectContext(ctx, "mysql", databaseURL)

	// handle error, if any
	if err != nil {
		log.Fatal("mysql failed to start")
	}

	// test database connection
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("could not ping the database")
	}

	// return database object
	return db
}
