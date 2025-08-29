package main

import (
	"database/sql"

	db2 "github.com/brinobruno/ports-and-adapters/adapters/db"
	"github.com/brinobruno/ports-and-adapters/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product example", 30)

	productService.Enable(product)
}
