package main

import (
	activityRepos "activities/domain/activity/repository"
	activityMysql "activities/domain/activity/repository/mysql"
	"activities/rest"
	"activities/service"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

// https://github.com/jsak2323/activities.git
var projectFolder = flag.String("folder", "./", "absolute path of project folder")

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ListenAddress := ":8000"
	e := echo.New()
	db := initDB("storage.db")
	migrate(db)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	var activityRepo activityRepos.Repository
	activityRepo = activityMysql.New(db)
	log.Println("Repository activities successfully initialized")

	var service = service.New(activityRepo)
	log.Println("Service successfully initialized")

	rest.New(service).Routes(e)
	log.Println("Router successfully initialized")

	fmt.Printf("%s apps running on port: %s\n", time.Now().Format("2006-01-02 15:04:05"), ListenAddress)
	if ListenAddress != "" {
		e.HidePort = true
		e.HideBanner = true
	}
	// Start Server
	s := &http.Server{
		Addr: ListenAddress,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt)
	go func() {
		<-sigChan
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		db.Close()
		log.Println("Database Successfully Stopped")

		os.Exit(0)
	}()

	err := e.StartServer(s)
	if err != nil {
		e.Logger.Fatal(err)
	}

}

func initDB(filepath string) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sqlx.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS activities(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		act VARCHAR NOT NULL,
		status INTEGER DEFAULT 1
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}
