package controllers

import (
	"github.com/1ssk/admin/initializers"
	"github.com/1ssk/admin/models"
	"github.com/gin-gonic/gin"
)

func AddTicket(c *gin.Context) {

	var body struct {
		FirstName   string
		LastName    string
		Email       string
		TicketCount uint
	}

	c.Bind(&body)

	post := models.Ticket{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, TicketCount: body.TicketCount}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllTicket(c *gin.Context) {

	var ticket []models.Ticket
	initializers.DB.Find(&ticket)

	c.JSON(200, gin.H{
		"ticket": ticket,
	})

}

func GetUserTicket(c *gin.Context) {
	id := c.Param("id")

	var ticket models.Ticket
	initializers.DB.First(&ticket, id)

	c.JSON(200, gin.H{
		"ticket": ticket,
	})
}

func DeleteUserTicket(c *gin.Context) {
	id := c.Param("id")

	var ticket models.Ticket
	initializers.DB.Delete(&ticket, id)

	c.JSON(200, gin.H{
		"Delete": "delete good",
	})

}
