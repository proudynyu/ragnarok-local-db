package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"server/src/domain/model"
	external_api "server/src/infrastrutcture/external-api"
)

func main() {
	base_url, err := external_api.NewApiUrl("https://ragnapi.com/api/v1/re-newal", "/monsters/", "/items/")

	if err != nil {
		log.Fatal("Not possible to create the base url")
	}

	var initial_monsters_id = 1001
	var end_monsters_id = 1002
	// var end_monsters_id = 3896

	// var initial_items_id = 501
	// var end_items_id = 60053

	fmt.Printf("init downloading files\n")
	for id := initial_monsters_id; id <= end_monsters_id; id++ {
		url, err := base_url.CreateMonsterUrlRequest(strconv.Itoa(id))

		if err != nil || url == "" {
            log.Fatal("Cannot GET the monster_id: ", id)
            time.Sleep(300 * time.Millisecond)
			continue
		}

        monster, err := external_api.Fetch[model.Monster](url)

        if err != nil && monster == "" {
            log.Fatal("Problem fetching monster_id: ", id)
        }

        fmt.Println(monster)

        time.Sleep(3 * time.Millisecond)
	}

}
