package routes

import (
	"net/http"
	"server/src/infrastrutcture/handlers"
)

func MonsterRouter() {
    http.HandleFunc("/monsters", handlers.MonsterHandler)
}
