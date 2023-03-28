package auth

import "project-template/model/entity"

func (c *AuthController) RegisterNewUser(newUser *entity.User) (*entity.User, error) {
	return c.repo.InsertNewUser(newUser)
}
