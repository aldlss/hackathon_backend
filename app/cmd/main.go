package main

import (
	"github.com/aldlss/hackathon_backend/app/cmd/api"
	"github.com/aldlss/hackathon_backend/app/cmd/dao"
)

func main() {
	dao.Init()
	api.Start()
}
