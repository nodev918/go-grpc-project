package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nodev918/go-grpc-api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
func Login(ctx *gin.Context, c pb.AuthServiceClient){
	b := LoginRequestBody{}
	
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}