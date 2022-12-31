package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/handlers"
	"github.com/knuls/horus/logger"
	"github.com/knuls/horus/middlewares"
	"github.com/knuls/horus/validator"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type config struct {
	Service  serviceConfig
	Store    storeConfig
	Server   serverConfig
	Security securityConfig
}

type serviceConfig struct {
	Name string
	Port int
}

type storeConfig struct {
	Client  string
	Host    string
	Port    int
	Name    string
	Timeout time.Duration
}

type serverConfig struct {
	Timeout struct {
		Read     time.Duration
		Write    time.Duration
		Idle     time.Duration
		Shutdown time.Duration
	}
}

type securityConfig struct {
	Allowed struct {
		Origins []string
		Methods []string
		Headers []string
	}
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
	c := viper.New()
	c.AddConfigPath(".")
	c.SetConfigName("config")
	c.SetConfigType("yaml")
	c.SetEnvPrefix("bennu")
	c.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	c.BindEnv("service.name")
	c.BindEnv("service.port")
	c.BindEnv("store.client")
	c.BindEnv("store.host")
	c.BindEnv("store.port")
	c.BindEnv("store.timeout")
	c.BindEnv("store.name")
	c.BindEnv("server.timeout.read")
	c.BindEnv("server.timeout.write")
	c.BindEnv("server.timeout.idle")
	c.BindEnv("server.timeout.shutdown")
	c.BindEnv("security.allowed.origins")
	c.BindEnv("security.allowed.methods")
	c.BindEnv("security.allowed.headers")
	c.BindEnv("security.allowCredentials")
	c.AutomaticEnv()
	var cfg config
	if err := c.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Error("config file not found", "error", err)
		} else {
			log.Error("config file read", "error", err)
		}
		return
	}
	err = c.Unmarshal(&cfg)
	if err != nil {
		log.Error("config decode error", "error", err)
		return
	}

	// db
	dbCtx, cancel := context.WithTimeout(context.Background(), cfg.Store.Timeout*time.Second)
	defer cancel()
	uri := fmt.Sprintf("%s://%s:%d", cfg.Store.Client, cfg.Store.Host, cfg.Store.Port)
	client, err := mongo.Connect(dbCtx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Error("db connect", "error", err)
		return
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Error("db disconnect", "error", err)
			return
		}
	}()
	pingCtx, cancel := context.WithTimeout(context.Background(), cfg.Store.Timeout*time.Second)
	defer cancel()
	if err = client.Ping(pingCtx, readpref.Primary()); err != nil {
		log.Error("db ping error: %v", err)
		return
	}

	// mux
	mux := chi.NewRouter()

	// middlewares
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.Security.Allowed.Origins,
		AllowedMethods:   cfg.Security.Allowed.Methods,
		AllowedHeaders:   cfg.Security.Allowed.Headers,
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
		log.Error("validator new", "error", err)
		return
	}

	// factory
	db := client.Database(cfg.Store.Name)
	factory := dao.NewFactory(db, v)

	// handlers
	mux.Mount("/user", handlers.NewUserHandler(log, factory).Routes())
	mux.Mount("/organization", handlers.NewOrganizationHandler(log, factory).Routes())
	mux.Mount("/auth", handlers.NewAuthHandler(log, factory).Routes())

	// server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Service.Port),
		Handler:      mux,
		ErrorLog:     log.GetStdLogger(),
		ReadTimeout:  cfg.Server.Timeout.Read * time.Second,
		WriteTimeout: cfg.Server.Timeout.Write * time.Second,
		IdleTimeout:  cfg.Server.Timeout.Idle * time.Second,
	}

	// listen
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("listen and serve", "error", err)
			return
		}
	}()
	log.Infof("starting %s service on port: %d", cfg.Service.Name, cfg.Service.Port)

	// shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	sig := <-sigCh
	log.Infof("signal: %s", sig.String())

	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.Server.Timeout.Shutdown*time.Second)
	defer cancel()
	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		log.Error("shutdown", "error", err)
		return
	}
}
