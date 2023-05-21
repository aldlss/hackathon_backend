package api

import (
	"context"
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/aldlss/hackathon_backend/app/cmd/api/handle"
	"github.com/aldlss/hackathon_backend/app/cmd/dao"
	"github.com/aldlss/hackathon_backend/app/cmd/dao/db"
	"github.com/aldlss/hackathon_backend/app/pkg/constants"
	"github.com/aldlss/hackathon_backend/app/pkg/errno"
	"github.com/aldlss/hackathon_backend/app/pkg/pack"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type testApi struct {
	suite.Suite
	ctx     context.Context
	baseUrl string
}

func (s *testApi) TestUserApi() {
	cli, err := client.NewClient()
	s.NoError(err)

	var wg sync.WaitGroup

	patches := gomonkey.ApplyFunc(db.AddUser, func(ctx context.Context, username string, password string) error {
		switch username {
		case "aya":
			return nil
		case "satori":
			return errno.DatabaseErr
		default:
			return errno.UnclassifiedErr
		}
	})
	defer patches.Reset()

	status, _, err := cli.Post(s.ctx, nil,
		fmt.Sprintf("%s/xunya/user/register/?username=aya&password=shameimaru", s.baseUrl),
		nil)
	s.Equal(constants.OK, status)
	s.NoError(err)

	wg.Add(4)
	userErr := func(username string, password string) {
		status, body, err := cli.Post(s.ctx, nil,
			fmt.Sprintf("%s/xunya/user/register/?username=%s&password=%s", s.baseUrl, username, password),
			nil)
		s.Equal(constants.OK, status)
		s.NoError(err)

		resp := pack.BaseResp{}
		err = sonic.Unmarshal(body, &resp)
		s.NoError(err)
		s.NotZero(resp.StatusCode)

		wg.Done()
	}

	go userErr("koishi", "")
	go userErr("", "scarlet")
	go userErr("", "")
	go userErr("ayaa", "wwww")
	wg.Wait()

	patches1 := gomonkey.ApplyFunc(db.CheckUser, func(ctx context.Context, username string, password string) (string, error) {
		if username == "aya" && password == "shameimaru" {
			return "123", nil
		}
		return "", errno.DatabaseErr
	})
	defer patches1.Reset()

	status, body, err := cli.Post(s.ctx, nil,
		fmt.Sprintf("%s/xunya/user/login/?username=aya&password=shameimaru", s.baseUrl), nil)
	s.Equal(constants.OK, status)
	s.NoError(err)
	resp := handle.LoginResp{}
	err = sonic.Unmarshal(body, &resp)
	s.NoError(err)
	s.Zero(resp.StatusCode)
	s.Equal("123", resp.Token)

	wg.Add(5)
	loginErr := func(username string, password string) {
		status, body, err := cli.Post(s.ctx, nil,
			fmt.Sprintf("%s/xunya/user/login?username=%s&password=%s", s.baseUrl, username, password),
			nil)
		s.Equal(constants.OK, status)
		s.NoError(err)

		resp := handle.LoginResp{}
		err = sonic.Unmarshal(body, &resp)
		s.NoError(err)
		s.NotZero(resp.StatusCode)

		wg.Done()
	}
	go loginErr("satori", "komeiji")
	go loginErr("koishi", "")
	go loginErr("", "scarlet")
	go loginErr("", "")
	go loginErr("aya", "wwww")
	wg.Wait()

}

func (s *testApi) SetupSuite() {
	go Start()
	dao.Init()
	time.Sleep(time.Second)
	s.ctx = context.Background()
	s.baseUrl = "http://[::1]:9961"
}

func (s *testApi) TearDownSuite() {

}

func TestAPI(t *testing.T) {
	suite.Run(t, new(testApi))
}
