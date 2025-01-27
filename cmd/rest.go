package cmd

import (
	"auth-repo/authentication"
	"auth-repo/cache"
	"auth-repo/config"
	"auth-repo/email"
	"auth-repo/logger"
	"auth-repo/repo"
	"auth-repo/rest"
	"auth-repo/rest/handlers"
	"auth-repo/rest/utils"
	"log/slog"

	"github.com/spf13/cobra"
)

var serveRestCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the goland Auth service",
	RunE:  serveRest,
}

func serveRest(cmd *cobra.Command, args []string) error {
	conf := config.GetConfig()

	logger.SetupLogger(conf.ServiceName)
	utils.InitValidator()

	redisClient, err := cache.NewRedisClient(conf.RedisUrl, conf.RedisTlsMode)
	if err != nil {
		slog.Error("Failed to create redis client", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}

	db, err := repo.MigrateDB(conf)
	if err != nil {
		slog.Error("Unable to connect to database", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}

	defer db.Close()

	cache := cache.NewCache(redisClient)

	userInfoRepo := repo.NewUserInfoRepo(db)
	if userInfoRepo == nil {
		slog.Error("Unable to connect to userInfoRepo", logger.Extra(map[string]any{
			"error": err.Error(),
		}))
		return err
	}
	mail := conf.Mail
	emailRepo := email.NewEmailService(&mail)
	if emailRepo == nil {
		slog.Error("Unable to connect to emailRepo", map[string]interface{}{"error": err.Error()})
	}

	auth := authentication.NewService(userInfoRepo, emailRepo, cache)
	if auth == nil {
		slog.Error("Unable to create auth service", map[string]interface{}{"error": err.Error()})
	}

	handler := handlers.NewHandler(conf, auth)
	server := rest.NewServer(conf, handler)

	server.Start()
	server.Wg.Wait()

	return nil
}
