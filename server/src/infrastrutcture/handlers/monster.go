package handlers

import (
	"fmt"
	"net/http"
)

func GetMonsters(w http.ResponseWriter, r *http.Request) {

}

func GetMonsterById(w http.ResponseWriter, r *http.Request) {

}

func CreateOrUpdateMonsterById(w http.ResponseWriter, r *http.Request) {

}

func MonsterHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
        case http.MethodGet:
            fmt.Fprintln(w, "GET: Retrieve information")
        case http.MethodPost:
            fmt.Fprintln(w, "POST: Retrieve information")
    }
}
