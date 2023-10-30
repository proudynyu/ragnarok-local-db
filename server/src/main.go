package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"server/src/domain/model"
	external_api "server/src/infrastrutcture/external-api"
)

func main() {
	base_url, err := external_api.NewApiUrl("https://ragnapi.com/api/v1/re-newal", "/monsters/", "/items/")

	if err != nil {
		log.Fatal("Not possible to create the base url")
	}

    maxConcurrentRequest := 2

    urlChan := make(chan string)
    resultCh := make(chan string)

	fmt.Printf("init downloading files\n")
    urls, err := base_url.CreateMonsterUrlRequest()

    if err != nil {
        log.Fatal(err)
    }

    for _, url := range *urls {
        urlChan<-url
    }

    close(urlChan)

    var wg sync.WaitGroup

    for concurrency := 0; concurrency <= maxConcurrentRequest; concurrency++ {
        wg.Add(1)
        go external_api.Worker[model.Monster](urlChan, resultCh, &wg)
        time.Sleep(3 * time.Millisecond)
    }

    go func() {
        wg.Wait()
        close(resultCh)
    }()

    for msg := range resultCh {
        fmt.Println(msg)
    }
}
