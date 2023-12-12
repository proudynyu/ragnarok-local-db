package routes

import (
	"server/src/infrastrutcture/handlers"

	"github.com/gorilla/mux"
)

func MonsterRouter(r *mux.Router) {
    r.HandleFunc("/monsters", handlers.GetMonsters).
        Methods("GET")

    r.HandleFunc("/monsters/{id}", handlers.GetOrCreateOrUpdateMonsterById).
        Methods("GET", "POST", "PUT")
}
