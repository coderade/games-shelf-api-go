package config

type Config struct {
	Port   int
	Env    string
	Secret string
	Db     struct {
		Dsn string
	}
}

var conf Config

func SetConfig(cfg Config) {
	conf = cfg
}

func GetConfig() Config {
	return conf
}
