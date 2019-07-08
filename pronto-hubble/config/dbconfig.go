package config

const (
	Mongo = iota
)

type DbType int

type DbConfig struct {
	Host string
	Port int
	User string
	Password string
	Database string
	dbType DbType
}
