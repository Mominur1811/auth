package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

func loadConfig() error {
	exit := func(err error) {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		slog.Warn(".env not found, that's okay!")
	}

	viper.AutomaticEnv()

	config = &Config{
		Mode:               Mode(viper.GetString("MODE")),
		ServiceName:        viper.GetString("SERVICE_NAME"),
		HttpPort:           viper.GetInt("HTTP_PORT"),
		JwtSecret:          viper.GetString("JWT_SECRET"),
		EnableRedisTLSMode: viper.GetBool("ENABLE_REDIS_TLS_MODE"),
		DB: DBConfig{
			Host:                viper.GetString("READ_DB_HOST"),
			Port:                viper.GetInt("READ_DB_PORT"),
			Name:                viper.GetString("READ_DB_NAME"),
			User:                viper.GetString("READ_DB_USER"),
			Pass:                viper.GetString("READ_DB_PASS"),
			MaxIdleTimeInMinute: viper.GetInt("READ_DB_MAX_IDLE_TIME_IN_MINUTE"),
			EnableSSLMode:       viper.GetBool("READ_DB_ENABLE_SSL_MODE"),
		},
		HealthCheckRoute: viper.GetString("HEALTH_CHECK_ROUTE"),
		Mail: Email{
			SmtpHost:   viper.GetString("MAIL_SMTP_HOST"),
			SmtpPort:   viper.GetString("MAIL_SMTP_PORT"),
			SourceMail: viper.GetString("MAIL_SOURCE_MAIL"),
			AppPass:    viper.GetString("MAIL_APP_PASS"),
		},
		RedisUrl:     viper.GetString("REDIS_URL"),
		RedisTlsMode: viper.GetBool("REDIS_TLS_MODE"),
	}

	v := validator.New()
	if err := v.Struct(config); err != nil {
		exit(err)
	}

	return nil
}
