package main

import "github.com/krlosw9/cursosGo/go-db-gorm/storage"

func main() {
	driver := storage.MySQL
	storage.New(driver)

}
