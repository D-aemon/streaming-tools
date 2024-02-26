package env

var DBcfgs DBConfig

func init() {
	DBcfgs = GetDBConfigFromEnv()
}
