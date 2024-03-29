package services

import (
	"errors"
	"github.com/go_user_role/backend/configurations/db"
	"github.com/go_user_role/backend/schemas"
	"github.com/go_user_role/backend/serializers"
	"log"
	"strconv"
)

type IRoleService interface {
	CreateRoleService(role *serializers.Role) (bool, error)
	GetAllRolesService(limit string, offset string) (int64, []serializers.RoleResponse)
}

type RoleService struct{}

// CreateRoleService for creating Service
func (r *RoleService) CreateRoleService(role *serializers.Role) (bool, error) {
	err := db.DB.Table(schemas.Roles).Where("role_name = ?", role.RoleName).First(&role).Error
	if err == nil {
		return false, errors.New("role already exits")
	}
	err = db.DB.Table(schemas.Roles).Create(&role).Error
	if err == nil {
		return false, errors.New("role creation failed")
	}
	return true, nil
}

// GetAllRolesService GetAllRoles service
func (r *RoleService) GetAllRolesService(limit string, offset string) (int64, []serializers.RoleResponse) {
	var count int64

	if err := db.DB.Table("role_based_access.roles").Count(&count).Error; err != nil {
		// Handle the error, e.g., log it or return an error to the caller
		log.Fatal(err)
	}
	count = 10
	roleLimit, err := strconv.Atoi(limit)
	if err != nil || roleLimit == 0 {
		roleLimit = -1
	}
	roleOffset, err := strconv.Atoi(offset)
	if err != nil || roleOffset == 0 {
		roleOffset = -1
	}
	var roles []serializers.RoleResponse
	db.DB.Table(schemas.Roles).Limit(roleLimit).Offset(roleOffset).Find(&roles)

	return count, roles
}
