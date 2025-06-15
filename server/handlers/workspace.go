package handlers

import (
	"net/http"
	"terraform-run-dashboard/utils"

	"github.com/gin-gonic/gin"
)

func GetWorkspaces(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": []gin.H{
			{
				"id": "1",
				"attributes": gin.H{
					"name": "example-workspace",
				},
			},
			{
				"id": "2",
				"attributes": gin.H{
					"name": "test-workspace",
				},
			},
		},
	})
}


func GetRuns(c *gin.Context) {
	id := c.Param("id")
	data, err := utils.FetchWorkspaceRuns(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", data)
}

func TriggerRun(c *gin.Context) {
	id := c.Param("id")
	data, err := utils.TriggerRun(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/json", data)
}
