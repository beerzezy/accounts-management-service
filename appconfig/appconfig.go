package appconfig

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
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
	Expiration uint32
	SigningKey string
}
