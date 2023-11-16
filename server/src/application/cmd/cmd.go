package cmd

import (
	"database/sql"
	"fmt"
	"log"
	external_api "server/src/infrastrutcture/external-api"
	"server/src/infrastrutcture/repositories"
)

func Exec(*sql.DB) {
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
