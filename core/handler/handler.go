package handler

import (
	//	"encoding/json"
	//	"io"
	//	"io/ioutil"
	"net/http"

	"github.com/devectron/moonlight/core/rest"
	"github.com/gin-gonic/gin"
)

func GetAllRest(c *gin.Context) {
	c.JSON(http.StatusOK, rest.GetRests())
}
