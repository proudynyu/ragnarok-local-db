package server

import (
	"errors"
	"fmt"
	"net/http"
)

type ApiUrl struct {
    base string
    monsters string
    items string
}

func (a *ApiUrl) CreateMonstersUrlRequest(monster_id string) (string, error) {
    if monster_id == "" {
        return "", errors.New("No Monster ID was passed")
    }

    url := a.base + a.monsters + monster_id

    if url == monster_id {
        return "", errors.New("No Monster ID was passed") 
    }

    return url, nil
}

func NewApiUrl(base string, monsters string, items string) (*ApiUrl, error) {
    base_url := ApiUrl {
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

func main () {
    base_url, err := NewApiUrl("https://ragnapi.com/api/v1/re-newal", "/monsters/", "/items/")

    if err != nil {
        errors.New("Not possible to create the base url")
        return
    }

    var initial_monsters_id = 1001
    var end_monsters_id = 3896

    // var initial_items_id = 501
    // var end_items_id = 60053

    for id := initial_monsters_id; id <= end_monsters_id; id++ {
        url, err := base_url.CreateMonstersUrlRequest(fmt.Sprint(id))

        if err != nil || url == "" {
            errors.New("Cannot GET the monster_{id}")
            continue
        }

        http.Get(url)

    }

    fmt.Printf("init downloading files")
}
