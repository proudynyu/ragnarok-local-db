package external_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/src/domain/model"
)

type ExternalApiUrl struct {
	base     string
	monsters string
	items    string
}

func (a *ExternalApiUrl) CreateMonsterUrlRequest(monster_id string) (string, error) {
	if monster_id == "" {
		return "", errors.New("No Monster ID was passed")
	}

	url := a.base + a.monsters + monster_id

	if url == monster_id {
		return "", errors.New("No Monster ID was passed")
	}

	return url, nil
}

func (a *ExternalApiUrl) CreateItemUrlRequest(item_id string) (string, error) {
	if item_id == "" {
		return "", errors.New("No Monster ID was passed")
	}

	url := a.base + a.monsters + item_id

	if url == item_id {
		return "", errors.New("No Monster ID was passed")
	}

	return url, nil
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

func Fetch[K model.Monster | model.Item](url string) (*K, error) {
    resp, err := http.Get(url)

    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    fmt.Printf("Fetched %s - Status Code: %d\n", url, resp.StatusCode)

    var target K

    err = json.NewDecoder(resp.Body).Decode(&target)

    if err != nil {
        return nil, err
    }

    return &target, nil
}
