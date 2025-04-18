package main

import (
	"github.com/1ssk/admin/controllers"
	"github.com/1ssk/admin/initializers"
	"github.com/gin-gonic/gin"
)

func Init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	Init()

	r.POST("/addTicket", controllers.AddTicket)
	r.GET("/getAllTicket", controllers.GetAllTicket)
	r.GET("/getUserTicket/:id", controllers.GetUserTicket)
	r.DELETE("/deleteUserTicket/:id", controllers.DeleteUserTicket)

	r.Run()
}
