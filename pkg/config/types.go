package config

type Config struct {
	AppConfig AppConfig
	DBConfig  DBConfig
	JwtConfig JwtConfig
}

type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type JwtConfig struct {
	Secret string
	Issuer string
}

type AppConfig struct {
	AppEnv string
	Port   int
}
