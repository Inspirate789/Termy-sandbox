package main

import (
	"context"
	"fmt"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/server"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/adapters/storage/postgres_storage/migrations"
	"git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/internal/domain/services/storage"
	_ "git.iu7.bmstu.ru/sam20u745/testing-web-2023-labs/swagger"
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	jaegerExporter "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var (
	// Store a global trace provider variable to clear it before closing
	tracer *tracesdk.TracerProvider
)

func NewTracer(url, name string) error {
	// Create the Jaeger exporter
	exp, err := jaegerExporter.New(jaegerExporter.WithCollectorEndpoint(jaegerExporter.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tracer = tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
		)),
	)
	return nil
}

func exitServer(logger *log.Logger, srv *server.Server) {
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	err := srv.Stop()
	if err != nil {
		logger.Errorf("Server Shutdown: %v", err)
	}
	logger.Info("Server exited")
}

func NewLogger(w io.Writer) *log.Logger {
	level, err := log.ParseLevel(os.Getenv("BACKEND_LOGLEVEL"))
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New()
	logger.SetOutput(w)
	logger.SetLevel(level)
	formatter := runtime.Formatter{
		ChildFormatter: &log.TextFormatter{
			FullTimestamp: true,
		},
		Package: true,
		File:    true,
		Line:    true,
	}
	formatter.Line = true
	logger.SetFormatter(&formatter)

	return logger
}

// @title			Termy API
// @version		1.0
// @description	This is a Termy backend API.
// @contact.name	API Support
// @contact.email	andreysapozhkov535@gmail.com
// @license.name	MIT
// @license.url	https://mit-license.org/
// @host			localhost:8080
// @BasePath		/api/v1
// @Schemes		http
func main() {
	//w := influx.NewInfluxWriter()
	//err := w.Open()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer w.Close()
	logger := NewLogger(os.Stdout)

	db, err := sqlx.Connect(os.Getenv("POSTGRES_DRIVER_NAME"), os.Getenv("POSTGRES_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	migration := migrations.NewMigration(db)
	err = migration.UpProd()
	if err != nil {
		panic(err)
	}

	storageService := storage.NewStorageService(postgres_storage.NewPostgresStorage(db), logger, 5)

	port, err := strconv.Atoi(os.Getenv("BACKEND_PORT"))
	if err != nil {
		panic(err)
	}

	if strings.ToLower(os.Getenv("BACKEND_TRACING")) == "true" {
		err = NewTracer(os.Getenv("TRACER_URL"), os.Getenv("BACKEND_NAME"))
		if err != nil {
			panic(err)
		}
		defer func(tracer *tracesdk.TracerProvider, ctx context.Context) {
			err := tracer.Shutdown(ctx)
			if err != nil {
				panic(err)
			}
		}(tracer, context.Background())
	}

	srv, err := server.NewServer(port, storageService, logger, tracer)
	if err != nil {
		panic(err)
	}

	go func() {
		err = srv.Start()
		if err != nil && !errors.Is(http.ErrServerClosed, err) {
			logger.Error(fmt.Sprintf("listen: %s\n", err))
		}
	}()
	logger.Infof("Server started at port %s", os.Getenv("BACKEND_PORT"))

	exitServer(logger, srv)
}
