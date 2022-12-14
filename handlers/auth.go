package handlers

import (
	"github.com/bacheha/horus/logger"
	"github.com/bacheha/horus/validator"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthHandler struct {
	Logger    *logger.Logger
	Validator *validator.Validator
	DB        *mongo.Database
}

func (h *AuthHandler) Routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/", nil) // GET /auth
	return mux
}

func NewAuthHandler(logger *logger.Logger, validator *validator.Validator, db *mongo.Database) *AuthHandler {
	return &AuthHandler{
		Logger:    logger,
		Validator: validator,
		DB:        db,
	}
}
