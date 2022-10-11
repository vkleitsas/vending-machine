package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	dbname   = os.Getenv("DB_NAME")
	port, _  = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	SqlxConn *sqlx.DB
)

var createUserTable = `CREATE TABLE IF NOT EXISTS "user" (
	id serial,
	username text,
	password text,
	deposit integer DEFAULT '-',
	role text,
	CONSTRAINT user_pkey PRIMARY KEY (id)
    
    
);`

var createProductTable = `CREATE TABLE IF NOT EXISTS "user" (
		id serial,
		username text,
		password text,
		deposit integer DEFAULT '-',
		role text,
		CONSTRAINT user_pkey PRIMARY KEY (id);
		
		CREATE TABLE IF NOT EXISTS "product"
		(
		id serial,
		product_name text,
		amount_available integer,
		cost integer,
		seler_id integer,
		CONSTRAINT product_pkey PRIMARY KEY (id),
		CONSTRAINT product_user_fkey FOREIGN KEY (seler_id)
		REFERENCES public."user" (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
    
);`

func databaseClient() (*sqlx.DB, error) {

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//log.Printf("psqInfo: %s", psqlInfo)
	SqlxConn, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	SqlxConn.SetMaxIdleConns(5)
	SqlxConn.SetMaxOpenConns(65)
	SqlxConn.SetConnMaxLifetime(5 * time.Second)
	if err = SqlxConn.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	return SqlxConn, nil
}

func initDB() {
	client, err := databaseClient()
	if err != nil {
		log.Printf("Error %s when opening database\n", err)
		return
	}

	tx := client.MustBegin()
	client.MustExec(createUserTable)
	client.MustExec(createProductTable)

	tx.Commit()
}
