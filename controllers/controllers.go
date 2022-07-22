package controllers

import (
	"github.com/alekseyklimenko/go-proj-bootstrap/controllers/some"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	some.RegisterHandlers(r)
}
