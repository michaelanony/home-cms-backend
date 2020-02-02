package model

import (
	"time"
)

type HomeUser struct {
	UId int64 `json:"u_id" db:"u_id"`
	UName string `json:"u_name" db:"u_name"`
	UNickname string `json:"u_nickname" db:"u_nickname"`
	URegisterIp string `json:"u_register_ip" db:"u_register_ip"`
	UAdmin bool `json:"u_admin" db:"u_admin"`
	ULevel int64 `json:"u_level" db:"u_level"`
	UPassword string `json:"u_password" db:"u_password"`
	UCreateTime time.Time `json:"u_create_time" db:"u_create_time"`
	UModifyTime time.Time `json:"u_modify_time" db:"u_modify_time"`
}

