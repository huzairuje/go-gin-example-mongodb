package loan

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Routes(router *gin.Engine, db *mongo.Database) {
	loanHandler := NewHandler(db)
	router.POST("/loans", loanHandler.Create)
	router.GET("/loans", loanHandler.List)
}
