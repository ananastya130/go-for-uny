package main

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)


func (a *App) indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// /api/search?q=часть
func (a *App) searchHandler(c *gin.Context) {
	q := strings.TrimSpace(c.Query("q"))
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "q parameter required"})
		return
	}
	like := "%" + q + "%"

	rows, err := a.DB.Query(`
		SELECT b.id, n.name, l.name, g.name, b.album_count
		FROM bands b
		JOIN names n ON b.name_id = n.id
		JOIN leaders l ON b.leader_id = l.id
		JOIN genres g ON b.genre_id = g.id
		WHERE n.name ILIKE $1 OR l.name ILIKE $1 OR g.name ILIKE $1
		ORDER BY b.id
	`, like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	results := []Band{}
	for rows.Next() {
		var b Band
		if err := rows.Scan(&b.ID, &b.Name, &b.Leader, &b.Genre, &b.AlbumCount); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, b)
	}
	c.JSON(http.StatusOK, results)
}
