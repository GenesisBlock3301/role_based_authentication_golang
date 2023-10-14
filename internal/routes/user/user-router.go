package user

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/internal/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(userRouter *gin.RouterGroup) {
	userRouter.POST("/create/", controller.GetALLRoleController)
}