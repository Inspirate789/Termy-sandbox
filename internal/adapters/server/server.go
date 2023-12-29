package server

import (
	"fmt"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/server/middleware"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"net/http"
	"os"
	"time"
)

type Server struct {
	srv            *http.Server
	storageService StorageService
	logger         *log.Logger
}

func (s *Server) addStudentRoutes(rg *gin.RouterGroup) {
	rg.POST("/units", s.postUnits)
	rg.PATCH("/units", s.patchUnits)
	rg.GET("/units", s.getUnits)
	rg.PUT("/units/models", s.getUnitsByModels)
	rg.PUT("/units/properties", s.getUnitsByProperties)
	rg.DELETE("/units", s.deleteUnit)

	rg.GET("/models", s.getModels)

	rg.GET("/elements", s.getModelElements)

	rg.GET("/properties", s.getProperties)
	rg.PUT("/properties/unit", s.getPropertiesByUnit)
	rg.POST("/properties", s.postProperties)

	rg.GET("/layers", s.getLayers)
}

func (s *Server) addEducatorRoutes(rg *gin.RouterGroup) {
	rg.POST("/models", s.postModels)

	rg.POST("/elements", s.postElements)

	rg.POST("/layers", s.postLayer)
}

func (s *Server) addAdminRoutes(rg *gin.RouterGroup) {
	rg.POST("/users", s.postUser)
	rg.GET("/stat", s.getStat)
}

func (s *Server) setupJWTMiddleware(requiredRoles []string, getUser middleware.UserCheckFunc, tracer *tracesdk.TracerProvider) (*jwt.GinJWTMiddleware, error) {
	var authMiddleware *jwt.GinJWTMiddleware
	var err error
	if tracer != nil {
		authMiddleware, err = middleware.MakeGinJWTMiddlewareWithTracing(requiredRoles, getUser, tracer)
	} else {
		authMiddleware, err = middleware.MakeGinJWTMiddleware(requiredRoles, getUser)
	}
	if err != nil {
		return nil, err
	}
	err = authMiddleware.MiddlewareInit()
	if err != nil {
		return nil, err
	}
	return authMiddleware, nil
}

func (s *Server) setupHandlers(router *gin.RouterGroup, tracer *tracesdk.TracerProvider) error {
	authMiddleware, err := s.setupJWTMiddleware([]string{
		entities.StudentRole,
		entities.EducatorRole,
		entities.AdminRole,
	}, s.storageService.GetUser, tracer)
	if err != nil {
		return err
	}
	educatorMiddleware, err := s.setupJWTMiddleware([]string{
		entities.EducatorRole,
		entities.AdminRole,
	}, s.storageService.GetUser, tracer)
	if err != nil {
		return err
	}
	adminMiddleware, err := s.setupJWTMiddleware([]string{
		entities.AdminRole,
	}, s.storageService.GetUser, tracer)
	if err != nil {
		return err
	}

	router.POST("/login", s.login(authMiddleware.LoginHandler))
	authRg := router.Group("", authMiddleware.MiddlewareFunc())
	authRg.GET("/refresh", s.refresh(authMiddleware.RefreshHandler))
	authRg.DELETE("/logout", s.logout(authMiddleware.LogoutHandler))

	studentRg := authRg.Group("")
	s.addStudentRoutes(studentRg)

	educatorRg := authRg.Group("", educatorMiddleware.MiddlewareFunc())
	s.addEducatorRoutes(educatorRg)

	adminRg := authRg.Group("", adminMiddleware.MiddlewareFunc())
	s.addAdminRoutes(adminRg)

	return nil
}

func NewServer(port int, storageMgr StorageService, logger *log.Logger, tracer *tracesdk.TracerProvider) (*Server, error) {
	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	router.UseRawPath = true
	router.UnescapePathValues = false

	url := ginSwagger.URL(os.Getenv("SWAGGER_PATH"))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(gin.RecoveryWithWriter(logger.Out))
	router.Use(gin.LoggerWithWriter(logger.Out))
	router.Use(middleware.ErrorResponseWriter(logger))
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	s := Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		storageService: storageMgr,
		logger:         logger,
	}

	apiRG := router.Group(os.Getenv("BACKEND_API_PREFIX"))

	err := s.setupHandlers(apiRG, tracer)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *Server) Start() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	return s.srv.Close()
}
