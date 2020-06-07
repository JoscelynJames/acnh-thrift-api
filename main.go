package main

import (
	"github.com/JoscelynJames/acnh-thrift-api/pg"
	"github.com/JoscelynJames/acnh-thrift-api/pg/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := pg.ConnectDB()
	defer db.Close()

	user := r.Group("/user")
	{
		user.POST("", services.CreateUser)
		user.GET("/:id", services.FindUser)
		user.PATCH("/:id", services.UpdateUser)
	}

	r.Run()
}
