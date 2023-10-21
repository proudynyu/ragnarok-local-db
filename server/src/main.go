package server

import (
	"fmt"
	"log"
	"strconv"
	"time"

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

    ch := make(chan string)

	for id := initial_monsters_id; id <= end_monsters_id; id++ {
		url, err := base_url.CreateMonsterUrlRequest(strconv.Itoa(id))

		if err != nil || url == "" {
			log.Fatal("Cannot GET the monster_{id}")
            time.Sleep(300 * time.Millisecond)
			continue
		}

        external_api.Fetch(url, ch)

        time.Sleep(300 * time.Millisecond)
	}

	fmt.Printf("init downloading files")
}
