package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	pgHost := getenv("PGHOST", "localhost")
	pgPort := getenv("PGPORT", "5433")
	pgDB := getenv("PGDB", "musicdb")
	pgUser := getenv("PGUSER", "musicuser")
	pgPass := getenv("PGPASS", "m0")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPass, pgDB)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("db ping: %v", err)
	}

	app := &App{DB: db}

	r := gin.Default()
	r.SetFuncMap(template.FuncMap{})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/", app.indexHandler)
	r.GET("/api/search", app.searchHandler)

	port := getenv("PORT", "8080")
	log.Printf("Server started at http://localhost:%s", port)
	r.Run(":" + port)
}

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
