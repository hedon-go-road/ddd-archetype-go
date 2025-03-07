package apihttp

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/hedon-go-road/ddd-archetype-go/docs"
	"github.com/hedon-go-road/ddd-archetype-go/internal/middleware"
	"github.com/hedon-go-road/ddd-archetype-go/pkg/response"
)

// @title           Match Service Swagger API
// @version         1.0
// @description     This is the open api doc for match service

// @host      :5050
// @BasePath  /

func SetupHTTPServer() {
	server := &API{}
	r := server.setupRouter()
	srv := &http.Server{
		Addr:              ":5050",
		Handler:           r.Handler(),
		ReadHeaderTimeout: time.Minute,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error().Err(err).Msg("error occurs in http server")
	}
}

type API struct{}

func (api *API) setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.WithRequestAndTrace())

	mg := r.Group("/api/v1")
	mg.POST("/hello", api.Hello)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

// Hello godoc
// @Summary hello
// @Description say hello to the server
// @Tags hello service
// @Accept json
// @Produce json
// @Param x-request-id header string false "Request ID"
// @Param HelloReq body HelloReq true "Hello Api Request Body"
// @Success 200 {object} HelloRsp
// @Failure 200 {object} string
// @Router /api/v1/hello [post]
func (api *API) Hello(c *gin.Context) {
	var req HelloReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.GinParamError(c, err)
		return
	}
	response.GinSuccess(c, HelloRsp{
		Name:    req.Name,
		Address: req.Address,
		Age:     req.Age,
		Married: req.Married,
		Server:  docs.SwaggerInfo.Description,
	})
}
