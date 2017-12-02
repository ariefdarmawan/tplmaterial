package model

import (
	"github.com/eaciit/toolkit"
)

type User struct {
	ID       string `bson:"_id" json:"_id"`
	Name     string
	Email    string
	Password string

	includePasswordWhenSave bool
}

func (u *User) TableName() string {
	return "appusers"
}

func (u *User) GetID() ([]string, []interface{}) {
	return []string{"_id"},
		[]interface{}{u.ID}
}

func (u *User) PreSave() {
	if u.includePasswordWhenSave {
		u.Password = toolkit.MD5String(u.Password)
	}
}

func Auth(userid, password string) {

}
