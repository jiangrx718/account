package handlers

import (
	"account/gopkg/utils"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Handler struct {
	router *gin.Engine

	//svc    studio.StudioService
	tracer trace.Tracer
	meter  metric.Meter
}

func NewHandler(router *gin.Engine) utils.HttpServerHandler {
	h := &Handler{
		router: router,

		//svc:    studio.New(),
		tracer: otel.Tracer("studio"),
		meter:  otel.Meter("studio"),
	}

	h.RegisterRoutes()
	return h
}

type Pagination struct {
	Page int `json:"page" query:"page"`
	Size int `json:"size" query:"size"`
}
