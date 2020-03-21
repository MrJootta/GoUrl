package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MrJootta/GoUrl/internal/storage"
)

type mysql struct {
	database *sql.DB
}

func New(host, port, username, password, database string) (storage.Service, error) {
	dataSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		username,
		password,
		host,
		port,
		database,
	)

	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		log.Fatal("Error connecting mysql database")
	}

	return &mysql{database: db}, err
}

func (db *mysql) NewCode(code, url string) (string, error) {

}

func (db *mysql) GetURL(code string) (storage.UrlCode, error) {

}

func (db *mysql) NewVisit(code string) error {

}

func (db *mysql) CodeInfo(code string) ([]storage.CodeVisit, error) {

}

func (p *mysql) Close() error {
	return p.database.Close()
}
