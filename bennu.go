package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/knuls/bennu/handlers"
	"github.com/knuls/horus/logger"
	"github.com/knuls/horus/middlewares"
	"github.com/knuls/horus/validator"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Service  ServiceConfig
	Store    StoreConfig
	Server   ServerConfig
	Security SecurityConfig
}

type ServiceConfig struct {
	Name string
	Port int
}

type StoreConfig struct {
	Client  string
	Host    string
	Port    int
	Name    string
	Timeout time.Duration
}

type ServerConfig struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

type SecurityConfig struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

func main() {
	// logger
	log, err := logger.New()
	if err != nil {
		fmt.Printf("logger new error: %v", err)
		os.Exit(1)
	}
	defer log.GetLogger().Sync()

	// config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("bennu")
	viper.BindEnv("service.name", "BENNU_SERVICE_NAME")
	viper.BindEnv("service.port", "BENNU_SERVICE_PORT")
	viper.BindEnv("store.client", "BENNU_STORE_CLIENT")
	viper.BindEnv("store.host", "BENNU_STORE_HOST")
	viper.BindEnv("store.port", "BENNU_STORE_PORT")
	viper.BindEnv("store.timeout", "BENNU_STORE_TIMEOUT")
	viper.BindEnv("store.name", "BENNU_STORE_NAME")
	viper.BindEnv("auth.allowedOrigins", "BENNU_AUTH_ALLOWED_ORIGINS")
	viper.AutomaticEnv()
	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("config file not found error: %v", err)
		} else {
			log.Fatalf("config file read error: %v", err)
		}
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("config decode error: %v", err)
	}

	// db
	dbCtx, cancel := context.WithTimeout(context.Background(), cfg.Store.Timeout*time.Second)
	defer cancel()
	uri := fmt.Sprintf("%s://%s:%d", cfg.Store.Client, cfg.Store.Host, cfg.Store.Port)
	client, err := mongo.Connect(dbCtx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalf("db disconnect error: %v", err)
		}
	}()
	pingCtx, cancel := context.WithTimeout(context.Background(), cfg.Store.Timeout*time.Second)
	defer cancel()
	if err = client.Ping(pingCtx, readpref.Primary()); err != nil {
		log.Fatalf("db ping error: %v", err)
	}

	// mux
	mux := chi.NewRouter()

	// middlewares
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.Security.AllowedOrigins,
		AllowedMethods:   cfg.Security.AllowedMethods,
		AllowedHeaders:   cfg.Security.AllowedHeaders,
		AllowCredentials: cfg.Security.AllowCredentials,
	}))
	mux.Use(middlewares.JSON)
	mux.Use(middlewares.RealIP)
	mux.Use(middlewares.RequestID)
	mux.Use(middlewares.Recoverer)
	mux.Use(middlewares.Logger(log))

	// validator
	v, err := validator.New()
	if err != nil {
		log.Fatalf("validator new error: %s", err.Error())
	}

	// handlers
	db := client.Database(cfg.Store.Name)
	mux.Mount("/user", handlers.NewUserHandler(log, v, db).Routes())
	mux.Mount("/organization", handlers.NewOrganizationHandler(log, v, db).Routes())
	mux.Mount("/auth", handlers.NewAuthHandler(log, v, db).Routes())

	// server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Service.Port),
		Handler:      mux,
		ErrorLog:     log.GetStdLogger(),
		ReadTimeout:  cfg.Server.ReadTimeout * time.Second,
		WriteTimeout: cfg.Server.WriteTimeout * time.Second,
		IdleTimeout:  cfg.Server.IdleTimeout * time.Second,
	}

	// listen
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen and serve error: %s", err.Error())
		}
	}()
	log.Infof("starting %s service on port: %d", cfg.Service.Name, cfg.Service.Port)

	// shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, os.Kill)
	sig := <-sigCh
	log.Infof("signal: %s", sig.String())

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout*time.Second)
	defer cancel()
	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		log.Fatalf("shutdown error: %s", err.Error())
	}
}
