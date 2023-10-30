package external_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"server/src/domain/model"
	"sync"
)

type ExternalApiUrl struct {
	base     string
	monsters string
	items    string
}

func (a *ExternalApiUrl) CreateMonsterUrlRequest() (*[]string, error) {
	var initial_monsters_id = 1001
	var end_monsters_id = 1002
	// var end_monsters_id = 3896

    var urls = []string {}

    for id := initial_monsters_id; id <= end_monsters_id; id++ {
        url := a.base + a.monsters + fmt.Sprint(id)
        urls = append(urls, url)
    }

    if len(urls) <= 0 {
		return nil, errors.New("No Monster ID was passed")
    }

	return &urls, nil
}

func (a *ExternalApiUrl) CreateItemUrlRequest(item_id string) (*[]string, error) {
	var initial_items_id = 501
	var end_items_id = 60053
    var urls = []string {}

    for id := initial_items_id; id <= end_items_id; id++ {
        url := a.base + a.monsters + fmt.Sprint(id)
        urls = append(urls, url)
    }

    if len(urls) <= 0 {
		return nil, errors.New("No Monster ID was passed")
    }

	return &urls, nil
}

func NewApiUrl(base string, monsters string, items string) (*ExternalApiUrl, error) {
	base_url := ExternalApiUrl{
		base,
		monsters,
		items,
	}

	if base_url.base == "" {
		return nil, errors.New("No base url was passed")
	}

	if base_url.monsters == "" {
		return nil, errors.New("No monsters url was passed")
	}

	if base_url.items == "" {
		return nil, errors.New("No items url was passed")
	}

	return &base_url, nil
}

func JsonParse[K model.Monster | model.Item](body []byte) (string, error) {
    var target K

    err := json.Unmarshal(body, &target)

    if err != nil {
        return "", err
    }

    formatted, err := json.MarshalIndent(target, "", " ")
    
    if err != nil {
        return "", err
    }

    return string(formatted), nil
}

func Fetch[K model.Monster | model.Item](url string, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()

    resp, err := http.Get(url)

    if err != nil {
        ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
        return
    }

    defer resp.Body.Close()

    fmt.Printf("Fetched %s - Status Code: %d\n", url, resp.StatusCode)

    body, err := io.ReadAll(resp.Body)

    if err != nil {
        ch <- fmt.Sprintf("Error reading body from %s: %v", url, err)
        return
    }

    formatted, err := JsonParse[K](body)

    if err != nil {
        ch <- fmt.Sprintf("Error parsing json %s: %v", url, err)
        return
    }

    ch <- formatted
}

func Worker[K model.Monster | model.Item](urlCh <-chan string, resultCh chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()

    for url := range urlCh {
        Fetch[K](url, resultCh, wg)
    }
}
