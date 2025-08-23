package api

import (
	// "fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var todos = []Todo{
	{ID: 1, Title: "Learn Go", Completed: false, CreatedAt: time.Now()},
	{ID: 2, Title: "Build a web app", Completed: true, CreatedAt: time.Now()},
}

func WelcomeEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Todo API!"})
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func PostTodos(c *gin.Context){
	var todo Todo

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, todo)
	c.JSON(http.StatusCreated, todo)
}


// func albumsByArtist(name string) ([]Album, error) {
// 	rows, err := db.Query("SELECT id, title, artist, price FROM album WHERE artist = $1", name)
// 	if err != nil {
// 		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 	}
// 	defer rows.Close()

// 	var albums []Album
// 	for rows.Next() {
// 		var a Album
// 		if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
// 			return nil, err
// 		}
// 		albums = append(albums, a)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return albums, nil
// }

// // albumByID queries for the album with the specified ID.
// func albumByID(id int64) (Album, error) {
// 	var a Album
// 	row := db.QueryRow("SELECT * FROM album WHERE id = $1", id)
// 	if err := row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
// 		return a, err
// 	}
// 	return a, nil
// }

// func addAlbum(album Album) (int64, error) {
// 	var id int64
// 	err := db.QueryRow(
// 		"INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id",
// 		album.Title, album.Artist, album.Price,
// 	).Scan(&id)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return id, nil
// }
