package appconfig

type Config struct {
	Environment string
	Server      ServerConfig
	Database    DatabaseConfig
	JWT         JWTConfig
	GatewayHost string
}

type ServerConfig struct {
	Port    uint
	Context string
}

type DatabaseConfig struct {
	Host     string
	Port     uint
	DBName   string
	Username string
	Password string
	SSLMode  string
}

type JWTConfig struct {
	Expiration string
	SigningKey string
}
