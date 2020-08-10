package app

import (
	"database/sql"
)

// User ...
type User struct {
	ID          string
	Login       string
	Email       string
	Password    string
	PasswordEnc string
}

// NewUser ...
func NewUser(id string, login string, email string, password string, passwordEnc string) *User {
	s := &User{
		ID:          id,
		Login:       login,
		Email:       email,
		Password:    password,
		PasswordEnc: passwordEnc,
	}
	return s
}

//
//
//

// UserSet ...
type UserSet struct {
	db *sql.DB
}

// // Init ...
// //     c - capacity
// func (us *UserSet) Init() error {
// 	db, err := sql.Open("MySQL", "root:password@/dbname")
// 	if err != nil {
// 		return err
// 	}
// 	us.db = db
// 	return nil
// }

// // Append ...
// func (us *UserSet) Append(id string, login string) {
// 	us.users = append(us.users, *NewUser(id, login))
// }

// // AppendUser ...
// func (us *UserSet) AppendUser(usr User) {
// 	us.users = append(us.users, usr)
// }

// // GetUsers ...
// func (us *UserSet) GetUsers() []User {
// 	return us.users
// }

// // Size ...
// func (us *UserSet) Size() (uint32, uint32) {
// 	return us.connectedUsers, us.capacity
// }

// // FindByLogin ...
// func (us *UserSet) FindByLogin(login string) *User {
// 	for _, i := range us.users {
// 		if i.Login == login {
// 			return &i
// 		}
// 	}
// 	return nil
// }

// // FindByID ...
// func (us *UserSet) FindByID(id string) *User {
// 	for _, i := range us.users {
// 		if i.ID == id {
// 			return &i
// 		}
// 	}
// 	return nil
// }
