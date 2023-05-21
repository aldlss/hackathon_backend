package db

import (
	"context"
	"github.com/aldlss/hackathon_backend/app/cmd/dao/model"
)

func PushContent(ctx context.Context, title string, desc string, token int64) (int64, error) {
	session := getPgDbSession(ContentPgDb, ctx)

	content := model.Content{
		Title:  title,
		Desc:   desc,
		UserId: token,
	}

	res := session.Create(&content)

	if res.Error != nil {
		return 0, res.Error
	}

	return content.Id, nil
}

func GetContent(ctx context.Context, token int64) (*model.Content, error) {
	session := getPgDbSession(ContentPgDb, ctx)

	content := model.Content{}

	var num int64
	res := session.Count(&num)
	if res.Error != nil {
		return nil, res.Error
	}

	res = session.Offset(random.Intn(int(num))).Take(&content)

	if res.Error != nil {
		return nil, res.Error
	}

	return &content, nil
}
