package handlers

import (
	"net/http"
	"time"

	"github.com/bacheha/horus/logger"
	"github.com/bacheha/horus/validator"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Token struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Scope     string             `json:"scope" bson:"scope" validate:"required"`
	Token     string             `json:"token" bson:"token" validate:"required"`
	Active    bool               `json:"active" bson:"active" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
}

type AuthHandler struct {
	Logger    *logger.Logger
	Validator *validator.Validator
	DB        *mongo.Database
}

func (h *AuthHandler) Routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Post("/login", h.Login)                  // POST /auth/login
	mux.Post("/register", h.Register)            // POST /auth/register
	mux.Post("/verify-email", h.VerifyEmail)     // POST /auth/verify-email
	mux.Post("/reset-password", h.ResetPassword) // POST /auth/reset-password
	mux.Post("/logout", h.Logout)                // POST /auth/logout
	mux.Route("/token", func(mux chi.Router) {
		mux.Post("/refresh", h.TokenRefresh) // POST /auth/token/refresh
	})
	return mux
}

func (h *AuthHandler) Login(rw http.ResponseWriter, r *http.Request) {
	// find user by username
	// if no user -> invalid user/pass
	// compare pass
	// if invalid pass -> invalid user/pass
	// create access & refresh tokens
	// set access token in resp, set refresh token in cookie
}

func (h *AuthHandler) Register(rw http.ResponseWriter, r *http.Request) {
	// validate user
	// create user as pending
	// send verify email with token
}

func (h *AuthHandler) ResetPassword(rw http.ResponseWriter, r *http.Request) {
	// find user
	// send password-reset email with token
}

func (h *AuthHandler) VerifyEmail(rw http.ResponseWriter, r *http.Request) {
	// update user status to active
	// de-activate verify email token
}

func (h *AuthHandler) VerifyResetPassword(rw http.ResponseWriter, r *http.Request) {
	// update user password
	// de-activate reset password token
}

func (h *AuthHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	// de-activate refresh and access token(s)
}

func (h *AuthHandler) TokenRefresh(rw http.ResponseWriter, r *http.Request) {
	// check if refresh token is valid
	// if valid -> create & respond with access token (in resp) & refresh token (in cookie)
}

func NewAuthHandler(logger *logger.Logger, validator *validator.Validator, db *mongo.Database) *AuthHandler {
	return &AuthHandler{
		Logger:    logger,
		Validator: validator,
		DB:        db,
	}
}
