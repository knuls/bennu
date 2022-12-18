package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/horus/logger"
	"github.com/knuls/horus/res"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Scope     string             `json:"scope" bson:"scope" validate:"required"`
	Token     string             `json:"token" bson:"token" validate:"required"`
	Active    bool               `json:"active" bson:"active" validate:"required"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId" validate:"required,oid"`
}

type AuthHandler struct {
	Logger    *logger.Logger
	Validator *validator.Validator
	DB        *mongo.Database
}

func (h *AuthHandler) Routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/csrf", h.CSRF)                     // GET /auth/csrf
	mux.Post("/login", h.Login)                  // POST /auth/login
	mux.Post("/register", h.Register)            // POST /auth/register
	mux.Post("/reset-password", h.ResetPassword) // POST /auth/reset-password
	mux.Post("/logout", h.Logout)                // POST /auth/logout
	mux.Route("/verify", func(mux chi.Router) {
		mux.Post("/email", h.VerifyEmail)                  // POST /auth/verify/email
		mux.Post("/reset-password", h.VerifyResetPassword) // POST /auth/verify/reset-password
	})
	mux.Route("/token", func(mux chi.Router) {
		mux.Post("/refresh", h.TokenRefresh) // POST /auth/token/refresh
	})
	return mux
}

func (h *AuthHandler) CSRF(rw http.ResponseWriter, r *http.Request) {
	// generate csrf token for POST / PATCH requests
}

func (h *AuthHandler) Login(rw http.ResponseWriter, r *http.Request) {
	// find user by email
	// if no user -> invalid user/pass
	// compare pass
	// if invalid pass -> invalid user/pass
	// create access & refresh tokens
	// set access token in resp & refresh token in cookie
}

func (h *AuthHandler) Register(rw http.ResponseWriter, r *http.Request) {
	// decode
	user := &User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if err == io.EOF {
		render.Render(rw, r, res.ErrDecode(errors.New("empty body")))
		return
	}
	if err != nil {
		render.Render(rw, r, res.ErrDecode(err))
		return
	}

	// validate
	if err := h.Validator.ValidateStruct(user); err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// users
	collection := h.DB.Collection("users")

	// ensure email does not exist
	count, err := collection.CountDocuments(r.Context(), bson.M{"email": user.Email})
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	if count > 0 {
		render.Render(rw, r, res.ErrBadRequest(errors.New("email already exists")))
		return
	}

	// populate system user fields
	now := time.Now()
	user.Verified = false
	user.CreatedAt = now
	user.UpdatedAt = now

	// hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		render.Render(rw, r, res.Err(err, http.StatusInternalServerError))
		return
	}
	user.Password = string(bytes)

	// insert
	result, err := collection.InsertOne(r.Context(), user)
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// create token
	// send verify email with token

	// render
	render.Status(r, http.StatusCreated)
	render.Respond(rw, r, &res.JSON{"id": result.InsertedID})
}

func (h *AuthHandler) ResetPassword(rw http.ResponseWriter, r *http.Request) {
	// find user by email
	// send password-reset email with token
}

func (h *AuthHandler) VerifyEmail(rw http.ResponseWriter, r *http.Request) {
	// update user verified to true
	// de-activate verify email token
}

func (h *AuthHandler) VerifyResetPassword(rw http.ResponseWriter, r *http.Request) {
	// update user password
	// de-activate reset password token
}

func (h *AuthHandler) TokenRefresh(rw http.ResponseWriter, r *http.Request) {
	// check if refresh token is valid
	// if valid -> create & respond with access token (in resp) & refresh token (in cookie)
}

func (h *AuthHandler) Logout(rw http.ResponseWriter, r *http.Request) {
	// de-activate refresh and access token(s)
}

func NewAuthHandler(logger *logger.Logger, validator *validator.Validator, db *mongo.Database) *AuthHandler {
	return &AuthHandler{
		Logger:    logger,
		Validator: validator,
		DB:        db,
	}
}
