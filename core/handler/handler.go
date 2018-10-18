package handler

import (
	//	"encoding/json"
	//	"io"
	//	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hihebark/moonlight/core/rest"
)

func GetAllRest(c *gin.Context) {
	c.JSON(http.StatusOK, rest.GetRests())
}
