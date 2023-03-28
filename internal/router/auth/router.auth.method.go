package auth

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"project-template/model/dto"
	"project-template/model/entity"
	"strings"
)

type RegisterNewUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

func (r *AuthRouter) RegisterNewUser(c echo.Context) error {
	u := new(RegisterNewUserDTO)
	if err := c.Bind(u); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	res, err := r.controller.RegisterNewUser(&entity.User{
		Email:    u.Email,
		Password: u.Password,
		FullName: u.FullName,
	})

	if err != nil {
		if strings.ContainsAny(err.Error(), "duplicate") {
			return c.JSON(http.StatusBadRequest, dto.ResponseWrapper{
				Status: "400",
				Msg:    "Bad Request, Email not unique",
				Data:   nil,
			})
		}
		return c.String(http.StatusInternalServerError, "Bad request")
	}
	return c.JSON(http.StatusOK, dto.ResponseWrapper{
		Status: "200",
		Msg:    "OK",
		Data: dto.UserInformationDTO{
			Id:       res.Id,
			Email:    res.Email,
			FullName: res.FullName,
			Access:   res.Access,
		},
	})
}
