package controllers

import (
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/serializers"
	"github.com/GenesisBlock3301/role_based_access_boilerplate_go/backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleController struct {
	RoleService services.RoleService
}

func NewRoleController(RoleService services.RoleService) *RoleController {
	return &RoleController{
		RoleService: RoleService,
	}
}

func (u *RoleController) CreateRoleController(ctx *gin.Context) {
	var roleInput serializers.Role
	if err := ctx.ShouldBindJSON(&roleInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := u.RoleService.CreateRoleService(&roleInput)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Successfully role created!"})
}

func (u *RoleController) GetALLRoleController(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	roleCount, roles := u.RoleService.GetAllRolesService(limit, offset)
	ctx.JSON(200, gin.H{
		"totalRole": roleCount,
		"roles":     roles,
	})
}
