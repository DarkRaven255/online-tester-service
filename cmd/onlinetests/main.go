package main

import (
	"context"
	"log"
	"online-tests/app"
	"online-tests/config"
	"online-tests/delivery/http"
	"online-tests/domain/domainmodel"
	"online-tests/repository"
	"online-tests/service"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB

func main() {

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{"*"},
	}))
	// defer postgresDB.Close()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logrus.Infof("Received %s signal", <-c)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	er := repository.NewEntryRepository(postgresDB)

	es := service.NewTestService(er)
	a := app.NewApp(es)

	http.NewHandler(e, a)

	log.Fatal(e.Start(":" + config.Cfg.Port))
}

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v\n", err)
	}

	postgresDB = initPostgres()
}

func initPostgres() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Cfg.DbAccess), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return migrate(db)
}

func migrate(db *gorm.DB) *gorm.DB {
	// db.AutoMigrate(&model.Results{})
	// db.Model(&model.Results{}).AddForeignKey("test_id", "onlinetests.tests(id)", "CASCADE", "CASCADE")
	// db.Model(&model.Results{}).AddForeignKey("user_id", "onlinetests.users(id)", "CASCADE", "CASCADE")

	db.AutoMigrate(&domainmodel.Answer{})
	db.AutoMigrate(&domainmodel.Question{})
	db.AutoMigrate(&domainmodel.Test{})

	return db
}
