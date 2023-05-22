package handle

import (
	"context"
	"encoding/base64"
	"github.com/aldlss/hackathon_backend/app/cmd/dao/db"
	"github.com/aldlss/hackathon_backend/app/pkg/errno"
	"github.com/aldlss/hackathon_backend/app/pkg/pack"
	"github.com/cloudwego/hertz/pkg/app"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strconv"
)

type commonContent struct {
	Title string `json:"title" vd:"$!=''"`
	Desc  string `json:"desc" vd:"$!=''"`
	Img   string `json:"img" vd:"$!=''"`
}

type fullContent struct {
	CommonContent commonContent `json:"common_content"`
	Id            int64         `json:"id"`
}

type GetResp struct {
	pack.BaseResp
	fullContent `json:"content"`
}

type GetContentReq struct {
	token int64 `query:"token" vd:"$!=''"`
}

type ContentNewReq struct {
	GetContentReq
	CommonContent commonContent `json:"common_content"`
}

func loadImg(name string) (string, error) {
	imgPath := os.Getenv("IMG_PATH")
	if imgPath == "" {
		log.Error("IMG_PATH is empty")
		return "", errno.ServiceErr
	}

	file, err := os.Open(path.Join(imgPath, name))
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func saveImg(name string, data string) error {
	imgPath := os.Getenv("IMG_PATH")
	if imgPath == "" {
		log.Error("IMG_PATH is empty")
		return errno.ServiceErr
	}

	newData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = os.WriteFile(path.Join(imgPath, name), newData, 0644)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func ContentGet(ctx context.Context, c *app.RequestContext) {
	var req GetContentReq
	err := c.BindAndValidate(&req)
	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	content, err := db.GetContent(ctx, req.token)
	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	img, err := loadImg(strconv.FormatInt(content.Id, 10) + ".png")
	if err != nil {
		SendBaseResponse(c, err)
		return
	}

	resp := GetResp{
		BaseResp: *pack.BuildBaseResp(errno.Success),
		fullContent: fullContent{
			CommonContent: commonContent{
				Title: content.Title,
				Desc:  content.Desc,
				Img:   img,
			},
			Id: content.Id,
		},
	}

	SendResponse(c, resp)
}

func ContentNew(ctx context.Context, c *app.RequestContext) {
	req := ContentNewReq{}

	err := c.BindAndValidate(&req)
	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	id, err := db.PushContent(ctx, req.CommonContent.Title, req.CommonContent.Desc, req.token)
	if err != nil {
		log.Error(err.Error())
		SendBaseResponse(c, err)
		return
	}

	// 虽然也不一定是 png
	err = saveImg(strconv.FormatInt(id, 10)+".png", req.CommonContent.Img)
	if err != nil {
		SendBaseResponse(c, err)
		return
	}

	SendBaseResponse(c, errno.Success)
}
