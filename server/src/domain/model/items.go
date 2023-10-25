package model

type Item struct  {
    Id string
    Item_id string
    Name string
    Description string
    Monsters []*Monster
}

func (i *Item) New(item_id string, name string, description string, monsters []*Monster) error {
    return nil
}


