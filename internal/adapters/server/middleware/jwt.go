package middleware

import (
	"context"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/entities"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"net/http"
	"os"
	"slices"
	"time"
)

type LoginResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

type UserCheckFunc func(entities.AuthRequest) (*entities.User, error)

const (
	identityKey = "id"
	roleKey     = "role"
)

func MakeGinJWTMiddleware(requiredRoles []string, getUser UserCheckFunc) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "jwt",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) { // for login (step 1)
			var request entities.AuthRequest
			if err := c.ShouldBind(&request); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			return getUser(request)
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims { // for login (step 2)
			if v, ok := data.(*entities.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					roleKey:     v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} { // for auth (step 1)
			claims := jwt.ExtractClaims(c)
			return &entities.User{
				ID:   int(claims[identityKey].(float64)),
				Role: claims[roleKey].(string),
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool { // for auth (step 2)
			v, ok := data.(*entities.User)
			return ok && slices.Contains(requiredRoles, v.Role)
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, LoginResponse{
				Token:  token,
				Expire: expire.Format(time.RFC3339),
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, LoginResponse{
				Token:  token,
				Expire: expire.Format(time.RFC3339),
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.Status(code)
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}

func MakeGinJWTMiddlewareWithTracing(requiredRoles []string, getUser UserCheckFunc, tracer *tracesdk.TracerProvider) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "jwt",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		Authenticator: func(c *gin.Context) (interface{}, error) { // for login (step 1)
			_, span := tracer.Tracer(os.Getenv("BACKEND_NAME")).Start(context.Background(), "Login start")
			defer span.End()

			var request entities.AuthRequest
			if err := c.ShouldBind(&request); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}
			return getUser(request)
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims { // for login (step 2)
			_, span := tracer.Tracer(os.Getenv("BACKEND_NAME")).Start(context.Background(), "Login end")
			defer span.End()

			if v, ok := data.(*entities.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					roleKey:     v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} { // for auth (step 1)
			_, span := tracer.Tracer(os.Getenv("BACKEND_NAME")).Start(context.Background(), "Auth start")
			defer span.End()

			claims := jwt.ExtractClaims(c)

			return &entities.User{
				ID:   int(claims[identityKey].(float64)),
				Role: claims[roleKey].(string),
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool { // for auth (step 2)
			_, span := tracer.Tracer(os.Getenv("BACKEND_NAME")).Start(context.Background(), "Auth end")
			defer span.End()

			v, ok := data.(*entities.User)
			return ok && slices.Contains(requiredRoles, v.Role)
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			_, span := tracer.Tracer(os.Getenv("BACKEND_NAME")).Start(context.Background(), "Unauthorized response")
			defer span.End()

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			_, span := tracer.Tracer(os.Getenv("BACKEND_NAME")).Start(context.Background(), "Login response")
			defer span.End()

			c.JSON(http.StatusOK, LoginResponse{
				Token:  token,
				Expire: expire.Format(time.RFC3339),
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			_, span := tracer.Tracer(os.Getenv("BACKEND_NAME")).Start(context.Background(), "Refresh response")
			defer span.End()

			c.JSON(http.StatusOK, LoginResponse{
				Token:  token,
				Expire: expire.Format(time.RFC3339),
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			_, span := tracer.Tracer(os.Getenv("BACKEND_NAME")).Start(context.Background(), "Logout response")
			defer span.End()

			c.Status(code)
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
