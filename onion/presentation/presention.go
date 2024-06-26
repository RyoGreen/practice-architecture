package presentation

import (
	"onion-architecture/application"
	"onion-architecture/domain"
	"onion-architecture/presentation/request"
	"onion-architecture/repository"
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
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		job, err := handler.FindByID(idInt)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, job)
	})

	r.POST("/jobs", func(c *gin.Context) {
		var job *request.JobRequest
		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err := handler.Create(job)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, nil)
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
		c.JSON(200, nil)
	})

	r.DELETE("/jobs/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = handler.Delete(idInt)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, nil)
	})

	r.Run()
}
