package model

import "time"

type CityStatus struct {
	CityName string `json:"city_name" db:"city_name"`
	ConfirmNum int64 `json:"confirm_num" db:"confirm_num"`
	CureNum int64 `json:"cure_num" db:"cure_num"`
	DeathNum int64 `json:"death_num" db:"death_num"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}


type StaticNum struct {
	Id int64 `json:"id" db:"id"`
	CityNumAll int64 `json:"city_num_all" db:"city_num_all"`
	ConfirmNumAll int64 `json:"confirm_num_all" db:"confirm_num_all"`
	DeathNumAll int64 `json:"death_num_all" db:"death_num_all"`
	CureNumAll int64 `json:"cure_num_all" db:"cure_num_all"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type CitySearchPy struct {
	Id int64 `json:"id" db:"id"`
	CityName string `json:"city_name" db:"city_name"`
	CityPy string `json:"city_py" db:"city_py"`
}