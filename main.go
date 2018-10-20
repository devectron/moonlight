package main

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"

	//"github.com/devectron/moonlight/core/handler"
	"github.com/devectron/moonlight/core/rest"
	"github.com/gin-gonic/gin"
)

const (
	banner = "       _..._\n     .::::. `.\n    :::::::.  :\n    ::::::::  :\n    `::::::' .'\n      `'::'-'\n"
)

func main() {
	fmt.Printf("%s   Moonlight ~ 0.0.0\n", banner)
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})
	rests := []rest.Rest{}
	rests = append(rests, rest.Rest{len(rests), "Luxure Restaurants", "https://www.telegraph.co.uk/content/dam/Travel/Destinations/Asia/China/Macau/Macau-restaurant-4-seasons.jpg", rest.Quality{0, 0, 0}})
	rests = append(rests, rest.Rest{len(rests), "Macau Restaurants", "https://djrcxf5ktozv0.cloudfront.net/data/venueimage/2184/image_w670_h440.jpg", rest.Quality{0, 0, 0}})
	rests = append(rests, rest.Rest{len(rests), "Unsukay Restaurants", "http://unsukay.com/wp-content/uploads/2017/02/MT-SMYRNA.jpg", rest.Quality{0, 0, 0}})
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, rests)
	})
	//r.POST("/rest/:id/vote", handler.Vote)
	//r.DELETE("/rest/:id", handlers.DeleteTodoHandler)
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, POST, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
