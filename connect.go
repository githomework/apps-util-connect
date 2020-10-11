package connect

import (
	"log"

	"github.com/SAP/go-hdb/driver"
	_ "github.com/denisenkom/go-mssqldb"
	secret "github.com/githomework/apps-util-secret"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-oci8"
)

// "<user name>/<password>@<server name>:1521/<service name>"
// Calling code should handle error.
func Oracle(key string, dsn string) (*sqlx.DB, error) {
	if key == "" {
		return sqlx.Open("oci8", dsn)
	}
	return sqlx.Open("oci8", secret.Decrypt(key, dsn))
}

func OracleMust(key string, dsn string) *sqlx.DB {
	if key == "" {
		db, err := sqlx.Open("oci8", dsn)
		if err != nil {
			log.Panicln(err)
		}
		return db
	}
	db, err := sqlx.Open("oci8", secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "hdb://<user name>:<password>@<server name>:<port>"
// Calling code should handle error.
func HANA(key string, dsn string) (*sqlx.DB, error) {
	if key == "" {
		return sqlx.Open(driver.DriverName, dsn)
	}
	return sqlx.Open(driver.DriverName, secret.Decrypt(key, dsn))
}

func HANAMust(key string, dsn string) *sqlx.DB {
	if key == "" {
		db, err := sqlx.Open(driver.DriverName, dsn)
		if err != nil {
			log.Panicln(err)
		}
		return db
	}
	db, err := sqlx.Open(driver.DriverName, secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "server=<server name>;user id=<user name>;password=<password>;port=1433;database=<database name>;encrypt=disable"
// Calling code should handle error.
func MSSQL(key string, dsn string) (*sqlx.DB, error) {
	if key == "" {
		return sqlx.Open("mssql", dsn)
	}
	db, err := sqlx.Open("mssql", secret.Decrypt(key, dsn))
	return db, err
}

func MSSQLMust(key string, dsn string) *sqlx.DB {
	if key == "" {
		db, err := sqlx.Open("mssql", dsn)
		if err != nil {
			log.Panicln(err)
		}
		return db
	}
	db, err := sqlx.Open("mssql", secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "host=<server name> user=<user name> dbname=<database name> password=<password> sslmode=disable port=5432"
// Calling code should handle error.
func Postgres(key string, dsn string) (*sqlx.DB, error) {
	if key == "" {
		return sqlx.Open("postgres", dsn)
	}
	return sqlx.Open("postgres", secret.Decrypt(key, dsn))
}

func PostgresMust(key string, dsn string) *sqlx.DB {
	if key == "" {
		db, err := sqlx.Open("postgres", dsn)
		if err != nil {
			log.Panicln(err)
		}
		return db
	}
	db, err := sqlx.Open("postgres", secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}
