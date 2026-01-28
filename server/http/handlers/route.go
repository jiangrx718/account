package handlers

import (
	"account/gopkg/log"
	"account/gopkg/utils"
	"account/server/http/handlers/tags"
	"account/server/http/middleware"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	otelgin "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (h *Handler) RegisterRoutes() {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "Authorization", middleware.OpenAppOpenKey)
	config.AllowAllOrigins = true
	h.router.Use(ginzap.RecoveryWithZap(log.Sugar().Desugar(), true))
	h.router.Use(cors.New(config))
	h.router.Use(otelgin.Middleware("account-api"))
	h.router.Use(sentrygin.New(sentrygin.Options{Repanic: utils.Debug()}))

	{
		g := h.router.Group("/api/account")
		tags.NewHandler(h.router).RegisterRoutes(g)
	}
}
