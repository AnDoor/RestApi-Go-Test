package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context,code int, msg string){

	ctx.JSON(code,gin.H{
		"menssage" : msg,
		"ErrorCode": code,
	})
} 

func sendSucces(ctx *gin.Context, operation string,data interface{}){
	ctx.Header("Content-type","application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler %s successfull",operation),
		"data": data,
	})
}