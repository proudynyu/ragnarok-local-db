package database

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func ConnectDb() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "https://localhost:7777")

    if err != nil {
        return nil, err
    }

    defer db.Close()

    err = db.Ping()

    if err != nil {
        return nil, err
    }

    return db, nil
}
