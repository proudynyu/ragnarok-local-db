package main

import (
    "server/src/application/cmd"
)

func main() {
    cmd.Exec()
    // dbConn, err := database.ConnectDb()
    //
    // if err != nil {
    //     log.Fatal("Could not connect to the Database")
    // }
}
