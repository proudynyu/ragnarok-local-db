package repositories

import (
	"database/sql"
	"encoding/json"
	"server/src/domain/model"
	"server/src/infrastrutcture/database"
)

type ItemsRepository struct {
    db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemsRepository {
    return &ItemsRepository{ db: db }
}

func (repo ItemsRepository) Create(api_response []byte) error {
	var target model.Item

	err := json.Unmarshal(api_response, &target)

	if err != nil {
		return err
	}

	m := model.Item{
        Id: target.Id,
        Item_id: target.Item_id,
        Name: target.Name,
        Description: target.Description,
	}

	tableName := "items"

	err = database.CreateRecord(repo.db, tableName, m)

	if err != nil {
		return err
	}

	return nil
}

func (repo ItemsRepository) Update(item_id string) error {
	return nil
}

func (repo ItemsRepository) FindByName(item_name string) error {
	return nil
}

func (repo ItemsRepository) FindById(item_id string) error {
	return nil
}
