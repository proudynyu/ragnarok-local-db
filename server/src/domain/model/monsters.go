package model

type Monster struct  {
    id string
    monster_id int16
    size string
    race string
    element_type string
    items []*Item
}

func (m *Monster) New(monster_id string, size string, race string, element_type string) error {
    return nil
}


