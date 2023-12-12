package external_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"server/src/domain/model"
	"server/src/infrastrutcture/repositories"
	"time"
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

func Fetch[K model.Monster | model.Item](url string) ([]byte, error) {
    resp, err := http.Get(url)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    fmt.Printf("Fetched %s - Status Code: %d\n", url, resp.StatusCode)

    body, err := io.ReadAll(resp.Body)

    if err != nil {
        return nil, err
    }

    return body, nil
}

func GetUrlsAndCreateRecord(urls *[]string) {
    repo := repositories.MonsterRepo

    for _, url := range *urls {
        response, err := Fetch[model.Monster](url)

        if err != nil {
            fmt.Printf("Cannot get the %v\n", url)
            time.Sleep(3 * time.Millisecond)
            continue
        }

        err = repo.Create(response)

        if err != nil {
            // add to the log
        }

        time.Sleep(3 * time.Millisecond)
    }
}
