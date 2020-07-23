package app

// Room ...
type Room struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
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
	Age         string `json:"age"`
	Profession  string `json:"profession"`
	Sex         string `json:"sex"`
	Orientation string `json:"orientation"`
	Hobby       string `json:"hobby"`
	Qualities   string `json:"qualities"`
	Health      string `json:"health"`
	ExtraInfo   string `json:"extraInfo"`
	Ability     string `json:"ability"`
	Phobia      string `json:"phobia"`
}
