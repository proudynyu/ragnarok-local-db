package repositories

import (
	"database/sql"
	"encoding/json"
	"server/src/domain/model"
	"server/src/infrastrutcture/database"
)

type MonsterRepository struct {
	db *sql.DB
}

func (repo MonsterRepository) Create(api_response []byte) error {
	var target model.Monster

	err := json.Unmarshal(api_response, &target)

	if err != nil {
		return err
	}

	m := model.Monster{
		Monster_id:   target.Monster_id,
		Size:         target.Size,
		Race:         target.Race,
		Element_type: target.Element_type,
		Name:         target.Name,
		Gif:          target.Gif,
	}

	tableName := "monsters"

	err = database.CreateRecord(repo.db, tableName, m)

	if err != nil {
		return err
	}

	return nil
}
func (repo MonsterRepository) Update(monster_id string) error {
	return nil
}

func (repo MonsterRepository) FindByName(monster_name string) error {
	return nil
}

func (repo MonsterRepository) FindById(monster_id string) error {
	return nil
}
