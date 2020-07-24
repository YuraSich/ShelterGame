package app

// User ...
type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

// NewUser ...
func NewUser(id string, login string) *User {
	s := &User{
		ID:    id,
		Login: login,
	}
	return s
}

//
//
//

// UserSet ...
type UserSet struct {
	users          []User
	connectedUsers uint32
	capacity       uint32
}

// Init ...
//     c - capacity
func (us *UserSet) Init(c uint32) {
	us.users = make([]User, 0)
	us.connectedUsers = 0
	us.capacity = c
}

// Append ...
func (us *UserSet) Append(id string, login string) {
	us.users = append(us.users, *NewUser(id, login))
}

// AppendUser ...
func (us *UserSet) AppendUser(usr User) {
	us.users = append(us.users, usr)
}

// GetUsers ...
func (us *UserSet) GetUsers() []User {
	return us.users
}

// Size ...
func (us *UserSet) Size() (uint32, uint32) {
	return us.connectedUsers, us.capacity
}

// FindByLogin ...
func (us *UserSet) FindByLogin(login string) *User {
	for _, i := range us.users {
		if i.Login == login {
			return &i
		}
	}
	return nil
}

// FindByID ...
func (us *UserSet) FindByID(id string) *User {
	for _, i := range us.users {
		if i.ID == id {
			return &i
		}
	}
	return nil
}
