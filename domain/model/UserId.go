package model

import "time"

type UserId struct {
	Id       string
	CreateAt time.Time
}

func (u *UserId) GenerateNew() {}
