package main

import (
	"fmt"
	"linked-page/config"
	"linked-page/db"
	"linked-page/grpcs"
	"linked-page/model"
	"linked-page/proto/page"
	"net"
	"sync"

	docs "linked-page/docs"
	"linked-page/router"
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

func SetupGRPCServer() {
	grpcServer := grpc.NewServer()
	pageGRPCService := grpcs.NewPageGRPC()
	page.RegisterPageServer(grpcServer, pageGRPCService)
	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}

	grpcServer.Serve(listener)
}

func main() {
	env := config.Env.ENV
	port := config.Env.PORT
	ssl := config.Env.SSL
	api_version := config.Env.API_VERSION

	if env == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Start the default gin server
	r := gin.Default()

	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	api := r.Group("/api")
	router.PageRoute(api)
	router.ListRoute(api)
	db.DB.AutoMigrate(&model.List{})
	db.DB.AutoMigrate(&model.Page{})

	//seed data
	// var pageModel = new(model.PageModel)
	// pageModel.SeedData()

	// var listModel = new(model.ListModel)
	// listModel.SeedData()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// start grpc server
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		SetupGRPCServer()
	}()

	if env == "LOCAL" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DeepLinking(true)))
	}

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, env, ssl, api_version)

	if ssl == "TRUE" {

		//Generated using sh generate-certificate.sh
		SSLKeys := &struct {
			CERT string
			KEY  string
		}{
			CERT: "./cert/myCA.cer",
			KEY:  "./cert/myCA.key",
		}

		r.RunTLS(":"+port, SSLKeys.CERT, SSLKeys.KEY)
	} else {
		r.Run(":" + port)
	}
}
