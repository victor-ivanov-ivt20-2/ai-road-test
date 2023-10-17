package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/auth"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/http/router"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/config"
	"github.com/victor-ivanov-ivt20-2/ai-road-test/internal/lib/logger/slogpretty"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	db := config.ConnectionDatabase(log, cfg)
	validate := validator.New()

	db.Table("users").AutoMigrate(&user.User{})

	userRepository := user.NewUserRepository(db)
	authService := auth.NewAuthService(cfg, userRepository, validate)

	authController := auth.NewAuthController(authService)
	userController := user.NewUsersController(userRepository)

	routes := router.NewRouter(cfg, log, userRepository, authController, userController)

	log.Info("starting AI Road test . . .", slog.String("env", cfg.Env))
	log.Debug("Debug mode is active")

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", cfg.HTTPServer.Port),
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Error("Server run error", slog.String("error", err.Error()))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.EnvLocal:
		log = setupPrettySlog()
	case config.EnvDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
