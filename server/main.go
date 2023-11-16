package main

import (
	"log"
	"server/src/application/cmd"
	"server/src/infrastrutcture/database"
)

func main() {
    dbConn, err := database.ConnectDb()

    if err != nil {
        log.Fatal("Could not connect to the Database")
    }

    cmd.Exec(dbConn)
}
