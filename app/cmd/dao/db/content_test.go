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

type testContent struct {
	suite.Suite
	ctx context.Context
}

func (s *testContent) TestGetContent() {
	content, err := GetContent(s.ctx, 5)
	s.NoError(err)
	s.Equal("Aya", content.Title)
	s.Equal("Aya is the best kawaii", content.Desc)
	s.Equal(int64(514), content.UserId)

	content, err = GetContent(s.ctx, 514)
	s.Error(err)
}

func (s *testContent) TestPushContent() {
	id, err := PushContent(s.ctx, "Satori", "Satori is the best kawaii", 114514)
	s.NoError(err)
	s.Equal(int64(2), id)
}

func (s *testContent) SetupSuite() {
	err := os.Setenv("CONTENT_TABLE_NAME", "test114514content")
	if err != nil {
		log.Fatal(err.Error())
	}

	initPgsql()

	s.ctx = context.Background()

	err = ContentPgDb.Migrator().AutoMigrate(&model.Content{})
	if err != nil {
		log.Fatal(err.Error())
	}

	err = ContentPgDb.Session(&gorm.Session{
		SkipHooks: true,
	}).Where("1=1").Delete(&model.Content{}).
		Create(&model.Content{
			Title:  "Aya",
			Desc:   "Aya is the best kawaii",
			UserId: 514,
		}).Error
	if err != nil {
		log.Fatal(err)
	}
}

func (s *testContent) TearDownSuite() {
	err := ContentPgDb.Migrator().DropTable(os.Getenv("CONTENT_TABLE_NAME"))
	if err != nil {
		log.Error(err.Error())
	}
}

func TestContentDb(t *testing.T) {
	suite.Run(t, new(testContent))
}
