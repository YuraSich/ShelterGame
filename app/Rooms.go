package app

import (
	"errors"
)

// Room ...
type Room struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity uint8  `json:"capacity"`
	Full     bool
	UsersIn  []User
}

// NewRoom ...
//     	id
//     	name
// 	    capacity
func NewRoom(id string, name string, capacity uint8) *Room {
	r := &Room{
		ID:       id,
		Name:     name,
		Capacity: capacity,
	}
	r.UsersIn = make([]User, 0)
	return r
}

// AddUser ...
func (r *Room) AddUser(user User) error {
	if r.Capacity == uint8(len(r.UsersIn)-1) {
		return errors.New("Room is fill")
	}
	r.UsersIn = append(r.UsersIn, user)
	return nil
}
