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
func Oracle(dsn string) (*sqlx.DB, error) {
	return open("oci8", "", dsn)
}

func OracleWithKey(key string, dsn string) (*sqlx.DB, error) {
	return open("oci8", key, dsn)
}

func OracleMust(dsn string) *sqlx.DB {
	return openMust("oci8", "", dsn)
}

func OracleMustWithKey(key string, dsn string) *sqlx.DB {
	return openMust("oci8", key, dsn)
}

// "hdb://<user name>:<password>@<server name>:<port>"
// Calling code should handle error.
func HANA(dsn string) (*sqlx.DB, error) {
	return open(driver.DriverName, "", dsn)
}

func HANAWithKey(key string, dsn string) (*sqlx.DB, error) {
	return open(driver.DriverName, key, dsn)
}
func HANAMust(dsn string) *sqlx.DB {
	return openMust(driver.DriverName, "", dsn)
}
func HANAMustWithKey(key string, dsn string) *sqlx.DB {
	return openMust(driver.DriverName, key, dsn)
}

// "server=<server name>;user id=<user name>;password=<password>;port=1433;database=<database name>;encrypt=disable"
// Calling code should handle error.
func MSSQL(dsn string) (*sqlx.DB, error) {
	return open("mssql", "", dsn)
}

func MSSQLWithKey(key string, dsn string) (*sqlx.DB, error) {
	return open("mssql", key, dsn)
}

func MSSQLMust(dsn string) *sqlx.DB {
	return openMust("mssql", "", dsn)
}

func MSSQLMustWithKey(key string, dsn string) *sqlx.DB {
	return openMust("mssql", key, dsn)
}

// "host=<server name> user=<user name> dbname=<database name> password=<password> sslmode=disable port=5432"
// Calling code should handle error.
func Postgres(dsn string) (*sqlx.DB, error) {
	return open("postgres", "", dsn)
}

func PostgresWithKey(key string, dsn string) (*sqlx.DB, error) {
	return open("postgres", key, dsn)
}

func PostgresMust(dsn string) *sqlx.DB {
	return openMust("postgres", "", dsn)
}

func PostgresMustWithKey(key string, dsn string) *sqlx.DB {
	return openMust("postgres", key, dsn)
}

func open(driverName string, key, dsn string) (*sqlx.DB, error) {
	if key == "" {
		return sqlx.Open(driverName, dsn)
	}
	return sqlx.Open(driverName, secret.Decrypt(key, dsn))
}

func openMust(driverName string, key, dsn string) *sqlx.DB {
	if key == "" {
		db, err := sqlx.Open(driverName, dsn)
		if err != nil {
			log.Panicln(err)
		}
		return db
	}
	db, err := sqlx.Open(driverName, secret.Decrypt(key, dsn))
	if err != nil {
		log.Panicln(err)
	}
	return db
}
