package routes

import (
	"server/src/infrastrutcture/handlers"

	"github.com/gorilla/mux"
)

func MonsterRouter(r *mux.Router) {
    r.HandleFunc("/monsters", handlers.GetMonsters).
        Methods("GET")

    r.HandleFunc("/monsters/{id}", handlers.GetMonsterById).
        Methods("GET")

    r.HandleFunc("/monsters/{id}", handlers.CreateMonsterById).
        Methods("POST")

    r.HandleFunc("/monsters/{id}", handlers.UpdateMonsterById).
        Methods("PUT")

    r.HandleFunc("/monsters/{id}", handlers.DeleteMonsterById).
        Methods("DELETE")
}
