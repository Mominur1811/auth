package config

import (
	"sync"
)

var confOnce = sync.Once{}

type Mode string

const DebugMode = Mode("debug")
const ReleaseMode = Mode("release")

type Email struct {
	SourceMail string `json:"source_mail" validate:"required,email"`
	AppPass    string `json:"app_pass" validate:"required"`
	SmtpHost   string `json:"smtp_host" validate:"required"`
	SmtpPort   string `json:"smtp_port" validate:"required"`
}

type DBConfig struct {
	Host                string `json:"host"                    validate:"required"`
	Port                int    `json:"port"                    validate:"required"`
	Name                string `json:"name"                    validate:"required"`
	User                string `json:"user"                    validate:"required"`
	Pass                string `json:"pass"                    validate:"required"`
	MaxIdleTimeInMinute int    `json:"max_idle_time_in_minute" validate:"required"`
	EnableSSLMode       bool   `json:"enable_ssl_mode"`
}

type Config struct {
	Mode               Mode     `json:"mode"                       validate:"required"`
	ServiceName        string   `json:"service_name"               validate:"required"`
	HttpPort           int      `json:"http_port"                         validate:"required"`
	JwtSecret          string   `json:"jwt_secret"                        validate:"required"`
	EnableRedisTLSMode bool     `json:"enable_redis_tls_mode"`
	DB                 DBConfig `json:"db_config"                         validate:"required"`
	HealthCheckRoute   string   `json:"health_check_route"                validate:"required"`
	Mail               Email    `json:"mail_config"                      validate:"required"`
	RedisUrl           string   `json:"redis_url"                        validate:"required"`
	RedisTlsMode       bool     `json:"redis_tls_mode"`
}

var config *Config

func GetConfig() *Config {
	confOnce.Do(func() {
		loadConfig()
	})

	return config
}
