package cmd

import (
	"fmt"
	"log"
	"net/http"
	// "server/src/infrastrutcture/database"
	"server/src/infrastrutcture/routes"
)

func Exec() {
    // err := database.ConnectDb()
    //
    // if err != nil {
    //     log.Fatal("Not possible to connect to the Database")
    // }

    r := routes.Router()

    server := &http.Server{
        Handler: r,
        Addr: ":8080",
    }

    fmt.Printf("Server is listening on port 8080... \n")
    log.Fatal(server.ListenAndServe(), "Failed to initialize the server")
}
