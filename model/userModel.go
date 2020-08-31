package model

import (
	"time"
)

type HomeUser struct {
	Id int64 `json:"id" db:"id"`
	UName string `json:"u_name" db:"u_name"`
	UPassword string `json:"u_password" db:"u_password"`
	UMoney int64 `json:"u_money" db:"u_money"`
	URegisterIp string `json:"u_register_ip" db:"u_register_ip"`
	URole int64 `json:"u_role" db:"u_role"`
	UpdateUser string `json:"update_user" db:"update_user"`
	UCreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
