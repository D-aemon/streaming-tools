package env

import "os"

type DBConfig struct {
	User   string
	Passwd string
	IP     string
	Port   string
	DbName string
}

type LogConfig struct {
	Level string
}

type ServerConfig struct {
	HttpPort string
	GrpcPort string
}

type EnvConfigMap struct {
	DBConfig
	LogConfig
	ServerConfig
}

var DBcfgs DBConfig
var LogCfgs LogConfig
var ServerCfgs ServerConfig

func init() {
	DBcfgs.GetConfigFromEnv()
	LogCfgs.GeConfigFromEnv()
	ServerCfgs.GetConfigFromEnv()
}

func (c DBConfig) GetConfigFromEnv() {
	c.User = os.Getenv("DB_USER")
	c.Passwd = os.Getenv("DB_PASSWD")
	c.IP = os.Getenv("DB_IP")
	c.Port = os.Getenv("DB_PORT")
	c.DbName = os.Getenv("DB_DBNAME")
}

func (c LogConfig) GeConfigFromEnv() {
	c.Level = os.Getenv("LOG_LEVEL")
}

func (s ServerConfig) GetConfigFromEnv() {
	s.HttpPort = os.Getenv("HTTP_PORT")
	s.GrpcPort = os.Getenv("GRPC_PORT")
}
