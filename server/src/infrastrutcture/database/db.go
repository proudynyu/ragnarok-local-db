package database

import (
	"database/sql"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
    Db *sql.DB
)

func ConnectDb()  error {
	database, err := sql.Open("sqlite3", "http://localhost:7777")

	if err != nil {
		return err
	}

	defer Db.Close()

	err = Db.Ping()

	if err != nil {
		return err
	}

    Db = database

	return nil
}

func CreateRecord(db *sql.DB, tableName string, data interface{}) error {
	var columns []string
	var placeholders []string
	var values []interface{}

	dataType := reflect.TypeOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		columns = append(columns, field.Name)
		placeholders = append(placeholders, "?")
		values = append(values, reflect.ValueOf(data).Field(i).Interface())
	}

	// Construct the SQL query
	insertSQL := "INSERT INTO " + tableName + " (" + strings.Join(columns, ",") + ") VALUES (" + strings.Join(placeholders, ",") + ")"

	// Execute the SQL query
	_, err := db.Exec(insertSQL, values...)
	if err != nil {
		return err
	}

	return nil
}

func UpdateRecord(db *sql.DB, tableName string, data interface{}, id string) error {
	return nil
}
