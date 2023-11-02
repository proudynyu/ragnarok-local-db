package repositories

import (
	"database/sql"
	"server/src/domain/model"
)

type MonsterRepository struct {
    db *sql.DB
}

func (repo *MonsterRepository) Create(monster *model.Monster) error {
    return nil
}

func (repo *MonsterRepository) Update(monster_id string) error {
    return nil
}

func (repo *MonsterRepository) FindByName(monster_name string) error {
    return nil
}

func (repo *MonsterRepository) FindById(monster_id string) error {
    return nil
}
