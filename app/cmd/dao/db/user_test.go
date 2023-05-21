package db

import (
	"context"
	"github.com/aldlss/hackathon_backend/app/cmd/dao/model"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"testing"
)

type testUser struct {
	suite.Suite
	ctx context.Context
}

func (s *testUser) TestAddUser() {
	err := AddUser(s.ctx, "aldlss", "114514")
	s.NoError(err)

	err = AddUser(s.ctx, "aya", "aini")
	s.NoError(err)

	err = AddUser(s.ctx, "123satori", "aini")
	s.NoError(err)

	err = AddUser(s.ctx, "aldlss", "sb")
	s.Error(err)
}

func (s *testUser) TestCheckUser() {
	id, err := CheckUser(s.ctx, "saldlss", "sb")
	s.Error(err)

	id, err = CheckUser(s.ctx, "ayaldlss", "nani")
	s.Error(err)

	id, err = CheckUser(s.ctx, "aaldlss", "aya")
	s.Error(err)

	id, err = CheckUser(s.ctx, "ayaldlss", "aya")
	s.NoError(err)
	s.Equal("1", id)
}

func (s *testUser) SetupSuite() {
	err := os.Setenv("USER_TABLE_NAME", "test114514user")
	if err != nil {
		log.Fatal(err.Error())
	}

	initPgsql()

	s.ctx = context.Background()

	err = UserPgDb.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = UserPgDb.Session(&gorm.Session{
		SkipHooks: true,
	}).Where("1=1").Delete(&model.User{}).
		Create(&model.User{
			Username: "ayaldlss",
			Password: "aya",
		}).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (s *testUser) TearDownSuite() {
	err := UserPgDb.Migrator().DropTable(os.Getenv("TABLE_NAME"))
	if err != nil {
		log.Error(err.Error())
	}
}

func TestUserDb(t *testing.T) {
	suite.Run(t, new(testUser))
}
