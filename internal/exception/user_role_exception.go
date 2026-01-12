package exception

import (
	"go-server-starter/internal/i18n"
	"net/http"
)

var (
	UserRoleNotFound      = UserRole.New(http.StatusNotFound, "user role not found", i18n.ExcUserRoleNotFound)
	UserRoleAlreadyExists = UserRole.New(http.StatusBadRequest, "user role already exists", i18n.ExcUserRoleAlreadyExists)
)
