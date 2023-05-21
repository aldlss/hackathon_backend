package db

import (
	"context"
	"github.com/aldlss/hackathon_backend/app/cmd/dao/model"
	"strconv"
)

func CheckUser(ctx context.Context, username string, password string) (string, error) {
	session := getPgDbSession(UserPgDb, ctx)

	user := model.User{
		Username: username,
		Password: password,
	}

	res := session.Where("username = ? AND password = ?", username, password).First(&user)

	if res.Error != nil {
		return "", res.Error
	}

	return strconv.FormatInt(user.Id, 10), nil
}

func AddUser(ctx context.Context, username string, password string) error {
	session := getPgDbSession(UserPgDb, ctx)

	user := model.User{
		Username: username,
		Password: password,
	}

	res := session.Create(&user)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
