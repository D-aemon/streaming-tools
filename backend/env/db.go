package env

import "os"

type DBConfig struct {
	User   string
	Passwd string
	IP     string
	Port   string
	DbName string
}

func GetDBConfigFromEnv() DBConfig {
	return DBConfig{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWD"),
		IP:     os.Getenv("DB_IP"),
		Port:   os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_DBNAME"),
	}
}
