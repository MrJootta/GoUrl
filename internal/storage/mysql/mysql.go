package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MrJootta/GoUrl/internal/storage"
	_ "github.com/go-sql-driver/mysql"
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
		log.Fatalf("error connecting mysql database: %s", err.Error())
	}

	return &mysql{database: db}, err
}

func (db *mysql) NewCode(code, url string) (string, error) {
	result, err := db.database.Exec(fmt.Sprintf("INSERT INTO url VALUES('%s', '%s')", code, url))
	if err != nil {
		return "error", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "error", err
	}

	return strconv.FormatInt(id, 10), err
}

func (db *mysql) GetURL(code string) (storage.UrlCode, error) {
	var result storage.UrlCode

	rows := db.database.QueryRow(fmt.Sprintf("SELECT url FROM url WHERE code='%s' LIMIT 1", code))

	err := rows.Scan(&result.URL)
	if err != nil {
		return storage.UrlCode{}, err
	}

	return result, err
}

func (db *mysql) NewVisit(code string) error {
	_, err := db.database.Exec(fmt.Sprintf("INSERT INTO visits VALUES('%s', CURRENT_TIMESTAMP)", code))

	return err
}

func (db *mysql) CodeInfo(code string) ([]storage.CodeVisit, error) {
	now := time.Now()
	after := now.Add(-24 * time.Hour)

	result, err := db.database.Query(
		"SELECT * FROM visits WHERE code=? AND timestamp > ?",
		code,
		after,
	)

	if err != nil || result == nil {
		return []storage.CodeVisit{}, err
	}

	var returnResult []storage.CodeVisit

	for result.Next() {
		visit := storage.CodeVisit{}

		err := result.Scan(&visit.Code, &visit.Time)
		if err != nil {
			log.Fatal(err)
		}

		returnResult = append(returnResult, visit)
	}

	return returnResult, nil
}

func (db *mysql) Close() error {
	return db.database.Close()
}
