package main

import (
	"context"
	"fmt"
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

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
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
	defer postgresDB.Close()

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

	log.Fatal(e.Start(":8081"))
}

func init() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v\n", err)
	}

	postgresDB = initPostgres()
}

func initPostgres() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Cfg.DB.Host, config.Cfg.DB.Port, config.Cfg.DB.User, config.Cfg.DB.Password, config.Cfg.DB.Dbname)
	log.Printf("connecting to PostgreSQL: host=%s, port=%d, user=%s, db=%s, pass=XXXX  \n", config.Cfg.DB.Host, config.Cfg.DB.Port, config.Cfg.DB.User, config.Cfg.DB.Dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return migrate(db)
}

func migrate(db *gorm.DB) *gorm.DB {
	// db.AutoMigrate(&model.User{})

	// db.AutoMigrate(&model.Results{})
	// db.Model(&model.Results{}).AddForeignKey("test_id", "onlinetests.tests(id)", "CASCADE", "CASCADE")
	// db.Model(&model.Results{}).AddForeignKey("user_id", "onlinetests.users(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&domainmodel.Test{})
	// db.Model(&model.Test{}).AddForeignKey("user_id", "onlinetests.users(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&domainmodel.Question{})
	db.AutoMigrate(&domainmodel.Answer{})

	return db
}
