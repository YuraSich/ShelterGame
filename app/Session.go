package app

import (
	"errors"
	"time"

	"github.com/segmentio/ksuid"
)

// Session ...
type Session struct {
	user       *User
	lastAction time.Time
	ID         string
}

// StartSession ...
func (s *Session) StartSession(u *User) (*Session, error) {

	n := &Session{
		user:       u,
		lastAction: time.Now(),
		ID:         ksuid.New().String(),
	}
	if r := recover(); r != nil {
		return nil, errors.New("ID generating error")
	}
	return n, nil
}
