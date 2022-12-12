package handlers

import (
	"github.com/bachehah/horus/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	Logger   *logger.Logger
	Validate *validator.Validate
	Client   *mongo.Client
}

func (h *UserHandler) Routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/", nil)  // GET /user
	mux.Post("/", nil) // POST /user
	mux.Route("/{id}", func(mux chi.Router) {
		mux.Get("/", nil) // GET /user/:id
	})
	return mux
}

func NewUserHandler(logger *logger.Logger, validate *validator.Validate, client *mongo.Client) *UserHandler {
	return &UserHandler{
		Logger:   logger,
		Validate: validate,
		Client:   client,
	}
}
