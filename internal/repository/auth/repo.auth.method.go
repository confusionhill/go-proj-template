package auth

import (
	"github.com/labstack/gommon/log"
	"project-template/model/entity"
)

func (r *AuthRepository) GetUserInfoByEmail(email string) (*entity.User, error) {
	rows, err := r.Database.Query(queryGetUserInfoByEmail, email)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	user := &entity.User{}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Password, &user.FullName, &user.Access)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (r *AuthRepository) InsertNewUser(newUser *entity.User) (*entity.User, error) {
	rows, err := r.Database.Query(queryInsertNewUser, newUser.Email, newUser.Password, newUser.FullName)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&newUser.Id, &newUser.Access)
		if err != nil {
			return nil, err
		}
	}

	return newUser, nil
}
