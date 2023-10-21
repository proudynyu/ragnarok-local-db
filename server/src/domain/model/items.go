package model

type Item struct  {
    id string
    item_id string
    name string
    description string
    monsters []*Monster
}

func (i *Item) New(item_id string, name string, description string, monsters []*Monster) error {
    return nil
}


