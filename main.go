package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"path/filepath"

	"github.com/devectron/moonlight/core"
	"github.com/devectron/moonlight/core/rest"
	"github.com/gin-gonic/gin"
)

const (
	banner = "       _..._\n     .::::. `.\n    :::::::.  :\n    ::::::::  :\n    `::::::' .'\n      `'::'-'\n"
)

func init() {
	rand.Seed(10)
}
func main() {
	fmt.Printf("%s   Moonlight ~ 1.0.0\n", banner)
	r := gin.Default()
	r.Use(core.Middleware())
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
	rests = append(rests, rest.Rest{len(rests), "Luxure Restaurants", "https://www.telegraph.co.uk/content/dam/Travel/Destinations/Asia/China/Macau/Macau-restaurant-4-seasons.jpg", rest.Quality{rand.Intn(10), rand.Intn(10), rand.Intn(10)}})
	rests = append(rests, rest.Rest{len(rests), "Macau Restaurants", "https://djrcxf5ktozv0.cloudfront.net/data/venueimage/2184/image_w670_h440.jpg", rest.Quality{rand.Intn(10), rand.Intn(10), rand.Intn(10)}})
	rests = append(rests, rest.Rest{len(rests), "Unsukay Restaurants", "http://unsukay.com/wp-content/uploads/2017/02/MT-SMYRNA.jpg", rest.Quality{rand.Intn(10), rand.Intn(10), rand.Intn(10)}})
	rests = append(rests, rest.Rest{len(rests), "March Restaurants", "http://wyndhamgardenmanama.com/wp-content/uploads/ab_resized/ITALIAN-restaurant-img.jpg", rest.Quality{rand.Intn(10), rand.Intn(10), rand.Intn(10)}})

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, rests)
	})

	r.POST("/rest", func(c *gin.Context) {
		var rest rest.Rest
		c.BindJSON(&rest)
		rests[rest.ID] = rest
		fmt.Printf("updating id:%v\n", rest.ID)
		c.JSON(http.StatusOK, "Updating rest value")
	})
	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
