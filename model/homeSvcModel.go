package model


type HostsConfig struct {
	HostsName	string	`json:"hosts_name" db:"hosts_name"`
	HostsIp	string	`json:"hosts_ip" db:"hosts_ip"`
	HostsSvc	string	`json:"hosts_svc" db:"hosts_svc"`
}
