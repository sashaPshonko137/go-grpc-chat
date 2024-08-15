package user

import "time"

type User struct {
	Id int32
	Name string
	Time time.Time
}

type UserInfo struct {
	Name string
}