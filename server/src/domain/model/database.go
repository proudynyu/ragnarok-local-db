package model

type DatabaseInterface interface {
    Create(api_response []byte) error
    Update(monster_id string) error
    FindByName(monster_name string) error
    FindById(monster_id string) error
}
