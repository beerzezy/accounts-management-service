package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/beerzezy/accounts-management-service/appconfig"
	"github.com/beerzezy/accounts-management-service/routes"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	config := InitConfig()
	db := InitDB(config)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	routes.RegisterRoutes(e, db, config)

	go StartServer(e, config)
	WaitForGracefulShutdown(e)
}

func StartServer(e *echo.Echo, cfg appconfig.Config) {
	port := fmt.Sprintf(":%d", cfg.Server.Port)
	if err := e.Start(port); err != nil {
		e.Logger.Info("shutting down the server")
	}
}

func WaitForGracefulShutdown(e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func InitConfig() appconfig.Config {
	var c appconfig.Config
	viper.SetDefault("database.sslmode", "disable")
	viper.SetConfigName("config")
	viper.AddConfigPath("./env/")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		} else {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
	viper.AutomaticEnv()
	viper.Unmarshal(&c)
	return c
}

func InitDB(c appconfig.Config) *gorm.DB {
	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.DBName,
		c.Database.Username,
		c.Database.Password,
		c.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: sqlLogger,
		DryRun: false,
	})
	if err != nil {
		panic(fmt.Errorf("error connecting to DB: %v", err))
	}
	if err != nil {
		panic(fmt.Errorf("error connecting to DB: %v", err))
	}

	return db
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
