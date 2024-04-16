package storage

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/krlosw9/cursosGo/go-db/pkg/product"
	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		fmt.Println("Esto iniciaria el driver de Mysql, pero no cree su implementacion...")
		// NewMySQLDB()
	case Postgres:
		newPostgresDB()
	default:
		fmt.Println("Driver no configurado")
	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		var user string = "krlos"
		password := url.QueryEscape("//654321")
		var host string = "localhost:5432"
		var dbName string = "godb"
		var connectionString string = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbName)
		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("can't do ping: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

// Pool return a unique instance of db
func Pool() *sql.DB {
	return db
}

func stringToNull(s string) sql.NullString {
	null := sql.NullString{String: s}
	if null.String != "" {
		null.Valid = true
	}
	return null
}

func timeToNull(t time.Time) sql.NullTime {
	null := sql.NullTime{Time: t}
	if !null.Time.IsZero() {
		null.Valid = true
	}
	return null
}

// DAOProduct factory of product.Storage
func DAOProduct(d Driver) (product.Storage, error) {
	switch d {
	case Postgres:
		return newPsqlProduct(db), nil
	default:
		return nil, fmt.Errorf("Driver not implemented")
	}
}
