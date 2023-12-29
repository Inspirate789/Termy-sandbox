package main

import (
	"fmt"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/benchmark"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.UseRawPath = true
	router.UnescapePathValues = false

	router.Use(gin.RecoveryWithWriter(os.Stdout))
	router.Use(gin.LoggerWithWriter(os.Stdout))

	router.GET("/metrics", prometheusHandler())
	router.POST("/bench/sqlx", func(ctx *gin.Context) {
		nRaw := ctx.Query("n")
		n, err := strconv.ParseInt(nRaw, 10, 64)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		log.Println("Start benchmarking SQLX")
		res, err := benchmark.UsersSqlxRepositoryBenchmark(int(n))
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	})
	router.POST("/bench/gorm", func(ctx *gin.Context) {
		nRaw := ctx.Query("n")
		n, err := strconv.ParseInt(nRaw, 10, 64)
		if err != nil {
			_ = ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		log.Println("Start benchmarking GORM")
		res, err := benchmark.UsersGormRepositoryBenchmark(int(n))
		if err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	})

	s := http.Server{
		Addr:    fmt.Sprintf(":2112"),
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := s.ListenAndServe()
		if err != nil && !errors.Is(http.ErrServerClosed, err) {
			panic(err)
		}
	}()

	<-quit
	println("Shutdown Server ...")

	err := s.Close()
	if err != nil {
		panic(err)
	}
	println("Server exited")
}
