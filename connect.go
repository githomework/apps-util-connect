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
func Oracle(key string, dsn string) *sqlx.DB {
	db, err := sqlx.Open("oci8", secret.Decrypt(key, dsn))
	if err != nil {
		log.Println(err)
	}
	return db
}

func OracleMust(key string, dsn string) *sqlx.DB {
	db, err := sqlx.Open("oci8", secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "hdb://<user name>:<password>@<server name>:<port>"
func HANA(key string, dsn string) *sqlx.DB {
	db, err := sqlx.Open(driver.DriverName, secret.Decrypt(key, dsn))
	if err != nil {
		log.Println(err)
	}
	return db
}

func HANAMust(key string, dsn string) *sqlx.DB {
	db, err := sqlx.Open(driver.DriverName, secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "server=<server name>;user id=<user name>;password=<password>;port=1433;database=<database name>;encrypt=disable"
func MSSQL(key string, dsn string) *sqlx.DB {
	db, err := sqlx.Open("mssql", secret.Decrypt(key, dsn))
	if err != nil {
		log.Println(err)
	}
	return db
}

func MSSQLMust(key string, dsn string) *sqlx.DB {
	db, err := sqlx.Open("mssql", secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "host=<server name> user=<user name> dbname=<database name> password=<password> sslmode=disable port=5432"
func Postgres(key string, dsn string) *sqlx.DB {
	db, err := sqlx.Open("postgres", secret.Decrypt(key, dsn))
	if err != nil {
		log.Println(err)
	}
	return db
}

func PostgresMust(key string, dsn string) *sqlx.DB {
	db, err := sqlx.Open("postgres", secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}
