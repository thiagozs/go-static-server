package main

import (
	"context"
	"flag"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/logger/zerolog"
)

var (
	Version  = "0.0.1"
	HostPort = flag.String("hostport", "127.0.0.1:8080", "host and port to listen on")
)

func main() {
	flag.Parse()

	if *HostPort == "" {
		*HostPort = "127.0.0.1:8080"
	}

	logger := zerolog.New()
	logger.SetLevel(hlog.LevelInfo)
	hlog.SetLogger(logger)

	h := server.Default(server.WithHostPorts(*HostPort))

	h.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH"}, // Allowed request methods
		AllowHeaders:     []string{"Origin"},                                // Allowed request headers
		ExposeHeaders:    []string{"Content-Length"},                        // Request headers allowed in the upload_file
		AllowCredentials: true,                                              // Whether cookies are attached
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	h.LoadHTMLGlob("public/*")

	h.Static("/", "./assets")

	h.GET("/", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(200, "index.html", nil)
	})

	h.Spin()
}
