package handlers

import (
	"fmt"
    "net/http"
	// "log"

	// "server/src/infrastrutcture/database"
	// "server/src/infrastrutcture/external-api"
	// "server/src/infrastrutcture/repositories"
)

// func createExternalRoutes() {
// 	base_url, err := external_api.NewApiUrl("https://ragnapi.com/api/v1/re-newal", "/monsters/", "/items/")
//
// 	if err != nil {
// 		log.Fatal("Not possible to create the base url")
// 	}
//
//     urls, err := base_url.CreateMonsterUrlRequest()
//
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     repositories.MonsterRepo = repositories.NewMonsterRepo(database.Db)
//
//     external_api.GetUrlsAndCreateRecord(urls)
// }

func GetMonsters(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there from Monsters")
}

func GetOrCreateOrUpdateMonsterById(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there from Create or Update Monsters")
}
