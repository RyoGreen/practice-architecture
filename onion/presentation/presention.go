package presentation

import (
	"architecture/onion/application"
	"architecture/onion/domain"
	"architecture/onion/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Run(repo repository.JobRepository) {
	r := gin.Default()
	handler := application.NewJobApplication(repo)
	r.GET("/jobs", func(c *gin.Context) {
		jobs, err := handler.FindAll()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, jobs)
	})
	r.GET("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		job, err := handler.FindByID(idInt)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, job)
	})
	r.POST("/jobs", func(c *gin.Context) {
		var job domain.Job
		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := handler.Create(&job)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, job)
	})
	r.PUT("/jobs/:id", func(c *gin.Context) {
		var job domain.Job
		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := handler.Update(&job)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, job)
	})

	r.DELETE("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		err = handler.Delete(idInt)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, nil)
	})

	r.Run()
}
