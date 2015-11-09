package main

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/mysql"
	"log"
)

type check struct {
	ID      int
	HostUID int
	Command string
}

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func main() {
	db, err := sqlx.Connect("mysql", "user=serverstat dbname=serverstat sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	//args := []string{""}
	/*args := []string{"-c 2", "-i", "0.2", "-w 2", "bartbart333.tk"}
	debugExec("ping", args)*/
	//args = []string{"aasasfasf;;;nu.nl"}
	//debugExec("cat", args)
}
