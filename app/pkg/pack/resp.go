package pack

import "github.com/aldlss/hackathon_backend/app/pkg/errno"

type BaseResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func BuildBaseResp(err error) *BaseResp {
	return baseResp(errno.ConvertErr(err))
}

func baseResp(err errno.ErrNo) *BaseResp {
	return &BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
