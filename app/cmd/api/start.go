package api

import (
	"context"
	"github.com/aldlss/hackathon_backend/app/cmd/api/handle"
	"github.com/aldlss/hackathon_backend/app/pkg/constants"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Start() {
	r := server.Default(
		server.WithHostPorts("[::]:9961"),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(128*1024*1024),
	)

	xunyaGroup := r.Group("/xunya")

	userGroup := xunyaGroup.Group("/user")
	userGroup.POST("/login", handle.UserLogin)
	userGroup.POST("/register", handle.UserRegister)

	contentGroup := xunyaGroup.Group("/content")
	contentGroup.GET("/get", handle.ContentGet)
	contentGroup.POST("/new", handle.ContentNew)

	relationGroup := xunyaGroup.Group("/relation")
	relationGroup.GET("/great", handle.GreatAction)

	r.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.Status(constants.NotFound)
	})

	r.NoMethod(func(c context.Context, ctx *app.RequestContext) {
		ctx.Status(constants.MethodNotAllowed)
	})

	r.Spin()
}
