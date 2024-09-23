package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCrimeEvent(c *gin.Context) {
	c.JSON(
		http.StatusAccepted,
		gin.H{
			"message": "memanggil fungsi send crime event",
		})
}
