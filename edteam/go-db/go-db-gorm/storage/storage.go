package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	once   sync.Once
	err    error
	host   string = "localhost"
	user   string = "krlos"
	dbName string = "godb"
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
		//fmt.Println("Esto iniciaria el driver de Mysql, pero no cree su implementacion...")
		newMysqlDB()
	case Postgres:
		newPostgresDB()
	default:
		fmt.Println("Driver no configurado")
	}
}

func newPostgresDB() {
	once.Do(func() {
		var port string = "5432"
		//fmt.Println("La contraseña es: ", password)
		dsn := fmt.Sprintf("host=%s user=%s password=//654321 dbname=%s port=%s sslmode=disable TimeZone=America/Bogota", host, user, dbName, port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Conectado a postgres")
	})
}

func newMysqlDB() {
	once.Do(func() {
		dsn := fmt.Sprintf("%s://654321@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, host, dbName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Conectado a mySQL")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}
