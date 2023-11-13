package main

import (
	"fmt"
	"log"

	// "server/src/infrastrutcture/database"
	external_api "server/src/infrastrutcture/external-api"
	"server/src/infrastrutcture/repositories"
)

func main() {
    // dbConn, err := database.ConnectDb()
    //
    // if err != nil {
    //     log.Fatal("Could not connect to the Database")
    // }

	base_url, err := external_api.NewApiUrl("https://ragnapi.com/api/v1/re-newal", "/monsters/", "/items/")

	if err != nil {
		log.Fatal("Not possible to create the base url")
	}

	fmt.Printf("init downloading files\n")
    urls, err := base_url.CreateMonsterUrlRequest()

    if err != nil {
        log.Fatal(err)
    }

    var monsterRepo repositories.MonsterRepository

    external_api.GetUrlsAndCreateRecord(urls, monsterRepo)
}
