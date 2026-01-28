package utils

import (
	"context"
	"log"
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
)

var doc *redoc.Redoc
var toolDoc *redoc.Redoc

type HttpServerHandler interface {
	RegisterRoutes()
}

func NewDocsHandler(router *gin.Engine) HttpServerHandler {
	h := &DocsHandler{
		router: router,
	}
	h.RegisterRoutes()
	return h
}

type DocsHandler struct {
	router *gin.Engine
}

func (h *DocsHandler) RegisterRoutes() {
	if doc != nil {
		h.router.Use(ginredoc.New(*doc))
	}

	if toolDoc != nil {
		h.router.Use(ginredoc.New(*toolDoc))
	}
}

type HttpServer struct {
	http.Server
	router   *gin.Engine
	handlers []HttpServerHandler
}

func NewHttpServer(listen string) *HttpServer {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(RequestID())
	if Debug() {
		gin.SetMode(gin.DebugMode)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	srv := &HttpServer{
		router: r,
		Server: http.Server{
			Addr:    listen,
			Handler: r,
		},
		handlers: []HttpServerHandler{NewDocsHandler(r)},
	}

	return srv
}

func (s *HttpServer) RegisterHandler(funcs ...func(*gin.Engine) HttpServerHandler) {
	for _, fun := range funcs {
		s.handlers = append(s.handlers, fun(s.router))
	}
}

func (s *HttpServer) GracefulStart(ctx context.Context) {
	go func() {
		// service connections
		log.Printf("Server listen on %s\n", s.Addr)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(c); err != nil {
		log.Printf("Server Shutdown: %s\n", err)
	}
	log.Println("Server exiting")
}

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("x-request-id")
		if id == "" {
			id = uuid.New().String()
		}
		c.Set("x-request-id", id)

		c.Next()

		c.Header("x-request_id", id)
	}
}
