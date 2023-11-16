package routes

import (
	"net/http"
	"server/src/infrastrutcture/handlers"
)

func routes() {
    http.HandleFunc("/monsters", handlers.MonsterHandler)
}
