package db

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	// Postgres driver for production
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// SQLite3 driver for testing purpose
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/andodevel/clock_server/bootstrap"
	"github.com/andodevel/clock_server/constants"
	"github.com/andodevel/clock_server/db/seeds"
	"github.com/andodevel/clock_server/models"
)

// TODO: Performance - use connection pool

var db *gorm.DB

// Init ...
func Init() {
	// Init database
	initConnections()
	autoDropTables()
	autoCreateTables()
	autoMigrateTables()
}

func initConnections() {
	var adapter string
	adapter = bootstrap.Prop(constants.DBAdapter)
	switch adapter {
	case "postgre":
		createPostgresConnection()
	case "sqlite3":
		fallthrough
	default:
		createSqlite3Connection()
	}
}

func createPostgresConnection() {
	var connectionString string
	connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		bootstrap.Prop(constants.DBUsername), bootstrap.Prop(constants.DBPassword),
		bootstrap.Prop(constants.DBHost), bootstrap.Prop(constants.DBPort),
		bootstrap.Prop(constants.DBDatabase), bootstrap.Prop(constants.DBSSLMode))

	var err error
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.Exec("CREATE EXTENSION postgis")

	idleConns, err := strconv.Atoi(bootstrap.Prop(constants.DBIdleConnections))
	if err == nil {
		db.DB().SetMaxIdleConns(idleConns)
	} else {
		db.DB().SetMaxIdleConns(10)
	}

	openConns, err := strconv.Atoi(bootstrap.Prop(constants.DBOpenConnections))
	if err == nil {
		db.DB().SetMaxOpenConns(openConns)
	} else {
		db.DB().SetMaxOpenConns(10)
	}
}

func createSqlite3Connection() {
	var err error
	db, err = gorm.Open("sqlite3", bootstrap.Prop(constants.DBDatabase))
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)
}

// CurrentDBConn ...
func CurrentDBConn() *gorm.DB {
	if db != nil {
		return db
	}

	initConnections()
	return db
}

func autoCreateTables() {
	if !CurrentDBConn().HasTable(&models.User{}) {
		CurrentDBConn().CreateTable(&models.User{})
	}

	// Seeders
	if bootstrap.IsInDevMode() {
		seeds.SeedTestData(CurrentDBConn())
	} else {
		seeds.SeedData(CurrentDBConn())
	}
}

func autoMigrateTables() {
	CurrentDBConn().AutoMigrate(
		&models.User{})
}

func autoDropTables() {
	if bootstrap.IsInDevMode() {
		CurrentDBConn().DropTableIfExists(
			&models.User{})
	}
}
