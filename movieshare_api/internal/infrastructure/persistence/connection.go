package persistence

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // DBMS specified driver.
)

// connection is a singleton instance to contain connection.
var connection *gorm.DB

// getDBConnection creates a DB connection and return.
func getDBConnection() *gorm.DB {
	if connection == nil {
		var error error

		clearDefaultConfig()

		connection, error = gorm.Open(getDBConfig())
		if error != nil {
			panic("failed to connect to the database")
		}
	}

	return connection.New()
}

// closeDBConnection closes a DB connection.
func closeDBConnection() {
	if connection != nil {
		error := connection.Close()
		if error != nil {
			panic(error)
		}

		connection = nil
	}
}

// clearDefaultConfig clears default connection config in OS environment.
func clearDefaultConfig() {
	// Unset these because login with service file isn't supported by golang postgres library.
	_ = os.Unsetenv("PGSERVICE")
	_ = os.Unsetenv("PGSERVICEFILE")
	_ = os.Unsetenv("PGREALM")
}

// getDBConfig gets DB config.
func getDBConfig() (string, string) {
	DBMS := os.Getenv("DBMS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBSSLMode := "sslmode=disable"

	if len(DBHost) == 0 || len(DBPort) == 0 || len(DBUser) == 0 || len(DBName) == 0 {
		panic("database config is not specified")
	}

	return DBMS, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s", DBHost, DBPort, DBUser, DBPass, DBName, DBSSLMode)
}

// getSSLMode gets SSL mode.
func getSSLMode() string {
	if os.Getenv("GO_ENV") == "local" {
		return "sslmode=disable"
	}
	return ""
}
