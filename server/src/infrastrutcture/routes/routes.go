package routes

import "github.com/gorilla/mux"

func Router() (*mux.Router) {
    r := mux.NewRouter()

    // insert routes
    MonsterRouter(r)

    return r
}
