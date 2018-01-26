package data

import (
	"time"
)

type User struct {
	Id       int
	Uuid     string
	Name     string
	Email    string
	Password string
	CreateAt time.Time
}

type Session struct {
	Id       int
	Uuid     string
	Email    string
	UserId   int
	CreateAt time.Time
}

func (session *Session) User(user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, cratead_at From users WHERE id=$1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreateAt)
	return
}
