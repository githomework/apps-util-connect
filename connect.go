package connect

import (
	"log"

	"github.com/SAP/go-hdb/driver"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-oci8"
)

// "<user name>/<password>@<server name>:1521/<service name>"
func Oracle(dsn string) *sqlx.DB {
	db, err := sqlx.Open("oci8", dsn)
	if err != nil {
		log.Println(err)
	}
	return db
}

func OracleMust(dsn string) *sqlx.DB {
	db, err := sqlx.Open("oci8", dsn)
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "hdb://<user name>:<password>@<server name>:<port>"
func HANA(dsn string) *sqlx.DB {
	db, err := sqlx.Open(driver.DriverName, dsn)
	if err != nil {
		log.Println(err)
	}
	return db
}

func HANAMust(dsn string) *sqlx.DB {
	db, err := sqlx.Open(driver.DriverName, dsn)
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "server=<server name>;user id=<user name>;password=<password>;port=1433;database=<database name>;encrypt=disable"
func MSSQL(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mssql", dsn)
	if err != nil {
		log.Println(err)
	}
	return db
}

func MSSQLMust(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mssql", dsn)
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// "host=<server name> user=<user name> dbname=<database name> password=<password> sslmode=disable port=5432"
func Postgres(dsn string) *sqlx.DB {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Println(err)
	}
	return db
}

func PostgresMust(dsn string) *sqlx.DB {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Panicln(err)
	}
	return db
}
