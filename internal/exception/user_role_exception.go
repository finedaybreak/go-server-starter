package exception

import (
	"go-server-starter/internal/i18n"
	"net/http"
)

var (
	UserRoleNotFound      = New(http.StatusNotFound, 21001, "user role not found", i18n.ExcUserRoleNotFound)
	UserRoleAlreadyExists = New(http.StatusBadRequest, 21002, "user role already exists", i18n.ExcUserRoleAlreadyExists)
)
