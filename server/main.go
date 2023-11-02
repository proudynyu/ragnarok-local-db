package main

import (
	"fmt"
	"log"
	"time"

	"server/src/domain/model"
	external_api "server/src/infrastrutcture/external-api"
)

func main() {
	base_url, err := external_api.NewApiUrl("https://ragnapi.com/api/v1/re-newal", "/monsters/", "/items/")

	if err != nil {
		log.Fatal("Not possible to create the base url")
	}

	fmt.Printf("init downloading files\n")
    urls, err := base_url.CreateMonsterUrlRequest()

    if err != nil {
        log.Fatal(err)
    }

    for _, url := range *urls {
        response, err := external_api.Fetch[model.Monster](url)

        if err != nil {
            fmt.Printf("Cannot get the %v\n", url)
            time.Sleep(3 * time.Millisecond)
            continue
        }

        fmt.Println(response)
        time.Sleep(3 * time.Millisecond)
    }
}
