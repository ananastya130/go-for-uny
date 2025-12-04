package main

import "database/sql"

type Name struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Leader struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Band struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Leader     string `json:"leader"`
	Genre      string `json:"genre"`
	AlbumCount int    `json:"album_count"`
}

type App struct {
	DB *sql.DB
}
