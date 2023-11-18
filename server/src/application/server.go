package application

import (
	"fmt"
	"log"
	"net/http"
	"server/src/infrastrutcture/routes"
)

func server() {
    // set routes available
    routes.MonsterRouter()

    // Start listening
    port := 5555
    fmt.Printf("Server is listening on port %d... \n", port)

    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

    if err != nil {
        log.Fatal("Error trying to initialize the server...")
    }
}
