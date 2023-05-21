package handle

import (
	"context"
	"github.com/aldlss/hackathon_backend/app/cmd/dao/db"
	"github.com/aldlss/hackathon_backend/app/pkg/errno"
	"github.com/aldlss/hackathon_backend/app/pkg/pack"
	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/sirupsen/logrus"
)

type userReq struct {
	username string `query:"username" vd:"$!=''"`
	password string `query:"password" vd:"$!=''"`
}

type LoginResp struct {
	pack.BaseResp
	Token string `json:"token"`
}

func getIdPw(ctx context.Context, c *app.RequestContext) (*userReq, error) {
	var req userReq
	err := c.BindAndValidate(&req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func UserRegister(ctx context.Context, c *app.RequestContext) {

	req, err := getIdPw(ctx, c)

	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	err = db.AddUser(ctx, req.username, req.password)
	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	SendBaseResponse(c, errno.Success)
}

func UserLogin(ctx context.Context, c *app.RequestContext) {

	req, err := getIdPw(ctx, c)

	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
	}

	token, err := db.CheckUser(ctx, req.username, req.password)
	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	SendResponse(c, LoginResp{
		*pack.BuildBaseResp(errno.Success),
		token,
	})
}
