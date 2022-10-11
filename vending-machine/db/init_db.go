package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	"time"
)

type PostgresClient struct {
	client *sqlx.DB
}

func NewPostgresClient(c *sqlx.DB) *PostgresClient {
	return &PostgresClient{
		client: c,
	}
}

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
	deposit integer,
	role text,
	CONSTRAINT user_pkey PRIMARY KEY (id)
    
    
);`

// var deleteProductTable = `DROP TABLE "product"`
var createProductTable = `
		CREATE TABLE IF NOT EXISTS "product"
		(
		id serial,
		name text,
		amount integer,
		cost integer,
		seller integer,
		CONSTRAINT product_pkey PRIMARY KEY (id),
		CONSTRAINT product_user_fkey FOREIGN KEY (seller)
		REFERENCES public."user" (id) MATCH SIMPLE
		ON UPDATE NO ACTION
		ON DELETE CASCADE
    
);`

func NewDatabaseClient() (*sqlx.DB, error) {

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

func InitDB() {
	client, err := NewDatabaseClient()
	if err != nil {
		log.Printf("Error %s when opening database\n", err)
		return
	}

	tx := client.MustBegin()
	client.MustExec(createUserTable)
	//client.MustExec(deleteProductTable)
	client.MustExec(createProductTable)

	tx.Commit()
}

func (p *PostgresClient) GetTransaction() (*sqlx.Tx, error) {
	tx, err := p.client.Beginx()
	if err != nil {
		return nil, err
	}
	return tx, nil
}
func (p *PostgresClient) Rollback(tx *sqlx.Tx) error {
	return tx.Rollback()

}
func (p *PostgresClient) Commit(tx *sqlx.Tx) error {
	return tx.Commit()

}

type InitDBInterface interface {
	GetTransaction() (*sqlx.Tx, error)
	Rollback(tx *sqlx.Tx) error
	Commit(tx *sqlx.Tx) error
}
