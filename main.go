package main

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/hihebark/moonlight/core/handler"
	//	"github.com/hihebark/moonlight/core/rest"
)

const (
	banner = "       _..._\n     .::::. `.\n    :::::::.  :\n    ::::::::  :\n    `::::::' .'\n      `'::'-'\n"
)

func main() {
	fmt.Printf("%s   Moonlight ~ 0.0.0\n", banner)
	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

	//	rests := []rest.Rest{}
	//	rests = append(rests, *rest.NewRest("1"))
	//	rests = append(rests, *rest.NewRest("2"))
	//	rests = append(rests, *rest.NewRest("3"))
	//r.POST("/rest/:id/vote", handler.Vote)
	r.GET("/index", handler.GetAllRest)
	//r.DELETE("/rest/:id", handlers.DeleteTodoHandler)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
