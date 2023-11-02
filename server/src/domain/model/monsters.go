package model

type MonsterInterface interface {
    Create(monster *Monster) error
    Update(monster_id string) error
    FindByName(monster_name string) (*Monster, error)
    FindById(monster_id string) (*Monster, error)
}

type Monster struct  {
    id string 
    Monster_id int16 `json:"monster_id"`
    Size string `json:"size"`
    Race string `json:"race"`
    Element_type string `json:"type"`
    Name string `json:"monster_info"`
    // Items []*Item
}
