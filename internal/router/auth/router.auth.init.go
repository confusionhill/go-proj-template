package auth

import (
	ac "project-template/internal/controller/auth"
)

type AuthRouter struct {
	controller *ac.AuthController
}

func Init(c *ac.AuthController) *AuthRouter {
	r := &AuthRouter{controller: c}
	return r
}
