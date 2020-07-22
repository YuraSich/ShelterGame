package app

// User ...
type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

// Room ...
type Room struct {
	ID       string `json:"id"`
	Name     string
	Capacity int
}

// Player ...
type Player struct {
	user  *User
	cards *CardSet
	room  *Room
	alive bool
}

// CardSet ...
type CardSet struct {
	//Возраст
	Age string
	//Профессия
	Profession string
	//Пол
	Sex string
	//Ориентация
	Orientation string
	//Хобби
	Hobby string
	//Качества
	Qualities string
	//Здоровье
	Health string
	//Доп инфа
	ExtraInfo string
	//Способность
	
	//Фобия
	Phobia string

}
