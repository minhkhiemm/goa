package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/k0kubun/pp"
	_ "github.com/lib/pq"
	"github.com/minhkhiemm/goa1/db/sqlboiler-postgres/models"
)
//go:generate sqlboiler --wipe psql

func main() {
	db, err := sql.Open("postgres","dbname=meocon host=localhost user=admin password=123 sslmode=disable")
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected")

	// p:= &models.Pilot{
	// 	ID:2,
	// 	Name:"luan",
	// }
	// err = p.Insert(context.Background(),db,boil.Infer())
	// if err != nil {
	// 	panic(err)
	// }

	all, err := models.Pilots().Count(context.Background(),db)
	if err != nil {
		panic(err)
	}

	pp.Println(all)
}