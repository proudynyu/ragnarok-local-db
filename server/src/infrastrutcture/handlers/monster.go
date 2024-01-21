package handlers

import (
	"fmt"
    "net/http"
)

func GetMonsters(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there from Monsters")
}

func GetMonsterById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there from Create or Update Monsters")
}

func CreateMonsterById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there from Create or Update Monsters")
}

func UpdateMonsterById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from update")
}

func DeleteMonsterById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from update")
}
