package some

import (
	"github.com/alekseyklimenko/go-proj-bootstrap/models"
	"github.com/alekseyklimenko/go-proj-bootstrap/models/requests"
	"github.com/alekseyklimenko/go-proj-bootstrap/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func RegisterHandlers(r *gin.Engine) {
	r.GET("/some/list", someList)
	r.GET("/some/get", someGet)
	r.POST("/some/add", someAdd)
}

func someList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":   1,
		"name": "Name",
	})
}

func someGet(c *gin.Context) {
	someId, err := services.Validation.ValidateIdParam(c.Query("someId"), "someId")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := []map[string]any{
		{
			"id":      1,
			"some_id": someId,
			"name":    "Name",
		},
		{
			"id":      2,
			"some_id": someId,
			"name":    "Name",
		},
	}
	c.JSON(http.StatusOK, result)
}

func someAdd(c *gin.Context) {
	var formData requests.Some
	something := models.NewSome()
	if err := c.ShouldBindWith(&formData, binding.Form); err == nil {
		err = services.Some.CreateNew(something, formData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		services.Processing.QueueItem(*something)
		c.JSON(http.StatusOK, gin.H{"success": true})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": services.Validation.GetErrors(err, something)})
	}
}
