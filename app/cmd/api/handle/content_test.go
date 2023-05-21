package handle

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
	"path"
	"testing"
)

type testContent struct {
	suite.Suite
	ctx         context.Context
	reimuImgStr string
	ayaImgStr   string
}

func (s *testContent) TestLoadImg() {
	img, err := loadImg("aya.jpg")
	if err != nil {
		return
	}
	s.Equal(s.ayaImgStr, img)
}

func (s *testContent) TestSaveImg() {
	err := saveImg("reimu.png", s.reimuImgStr)
	s.NoError(err)
}

func (s *testContent) SetupSuite() {
	imgPath := "F:\\code\\competition\\NJUhackathon2023\\img\\test"
	err := os.Setenv("IMG_PATH", imgPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	s.ctx = context.Background()

	file, err := os.Open(path.Join(imgPath, "reimu.txt"))
	if err != nil {
		log.Fatal(err.Error())
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.Remove(path.Join(imgPath, "reimu.png"))
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err.Error())
	}

	s.reimuImgStr = string(data)

	file, err = os.Open(path.Join(imgPath, "aya.txt"))
	if err != nil {
		log.Fatal(err.Error())
	}

	data, err = io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	s.ayaImgStr = string(data)
}

func (s *testContent) TearDownSuite() {

}

func TestContentDb(t *testing.T) {
	suite.Run(t, new(testContent))
}
