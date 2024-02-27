package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// App config struct
type Config struct {
	Server      ServerConfig
	MySQL       MySQLConfig
	Logger      Logger
	Metrics     Metrics
	CDN         CDN
	ImgSvc      ImageService
	InternalAPI InternalAPIConfig
}

// Server config struct
type ServerConfig struct {
	AppVersion        string `env:"APP_VERSION"`
	Port              string `env:"PORT"`
	Mode              string `env:"MODE"`
	JwtSecretKey      string `env:"JWT_SECRET_KEY"`
	LofSecretKey      string `env:"LOF_SECRET_KEY"`
	ReadTimeout       int    `env:"READ_TIMEOUT"`
	WriteTimeout      int    `env:"WRITE_TIMEOUT"`
	CtxDefaultTimeout int    `env:"CTX_DEFAULT_TIMEOUT"`
	Debug             bool   `env:"DEBUG"`
}

// Metrics config
type Metrics struct {
	URL         string `env:"METRICS_URL"`
	ServiceName string `env:"METRICS_SERVICE_NAME"`
}

// Logger config
type Logger struct {
	Development       bool   `env:"LOGGER_DEVELOPMENT"`
	DisableCaller     bool   `env:"LOGGER_DISABLE_CALLER"`
	DisableStacktrace bool   `env:"LOGGER_DISABLE_STACKTRACE"`
	Encoding          string `env:"LOGGER_ENCODING"`
	Level             string `env:"LOGGER_LEVEL"`
}

type MySQLConfig struct {
	WriteURI        string   `env:"MYSQL_WRITE_URI"`
	ReadURIs        []string `env:"MYSQL_READ_URIS"`
	MaxIdleConns    int      `env:"MYSQL_MAX_IDLE_CONNS"`
	MaxOpenConns    int      `env:"MYSQL_MAX_OPEN_CONNS"`
	ConnMaxLifeTime int      `env:"MYSQL_CON_MAX_LIFE_TIME"`
	Timeout         int      `env:"MYSQL_TIMEOUT"`
	Debug           bool     `env:"MYSQL_DEBUG"`
}

type MySQLURI struct {
	MasterURI string
	SlaveURIs []string
}

type Database string

const (
	DefaultDB Database = "default"
	UserDB    Database = "userDB"
)

func (config MySQLConfig) GetURI(dn Database) (MySQLURI, error) {
	switch dn {
	case DefaultDB:
		return MySQLURI{
			MasterURI: config.WriteURI,
			SlaveURIs: config.ReadURIs,
		}, nil
	default:
		return MySQLURI{}, fmt.Errorf("not supported database")
	}

}

type CDN struct {
	CDN_URL string `env:"CDN_URL"`
}

type ImageService struct {
	Host      string `env:"IMG_SVC_HOST"`
	UploadKey string `env:"IMG_SVC_UPLOAD_KEY"`
	SecretKey string `env:"IMG_SVC_SECRET_KEY"`
}

// InternalAPIConfig is struct for internal api config
type InternalAPIConfig struct {
	BaseURL      string   `env:"INTERNAL_API_URL" envDefault:"http://localhost:8080"`
	APIKey       string   `env:"INTERNAL_API_KEY"`
	AcceptedKeys []string `env:"INTERNAL_API_ACCEPTED_KEYS" envDefault:"lof-lamviectot-api-key"`
	Source       string   `env:"INTERNAL_API_SOURCE" envDefault:"lof-usertest-api"`
}

// Load config file from given path
func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
