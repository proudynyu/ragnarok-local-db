package model

type Monster struct  {
    id string 
    Monster_id int16 `json:"monster_id"`
    Size string `json:"size"`
    Race string `json:"race"`
    Element_type string `json:"type"`
    // Items []*Item
}

func (m *Monster) New(monster_id string, size string, race string, element_type string) error {
    return nil
}
