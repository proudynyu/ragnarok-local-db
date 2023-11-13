package usecases

import (
	"server/src/domain/model"
)

type MonsterRepoInterface struct {
    db *model.MonsterInterface
}

func (m *MonsterRepoInterface) Create(api_response []byte) error {
    return nil
}

func (m MonsterRepoInterface) Update(monster_id string) error {
    return nil
}

func (m MonsterRepoInterface) FindByName(monster_name string) error {
    return nil
}

func (m MonsterRepoInterface) FindById(monster_id string) error {
    return nil
}
