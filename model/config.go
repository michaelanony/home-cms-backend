package model

type Config struct {
	DevMysqlDb string `toml:"DevMysqlDb"`
	DevRedis string `toml:"DevRedis"`
	Port	string `toml:"Port"`
}
