package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/horus/logger"
	"github.com/knuls/horus/middlewares"
	"github.com/knuls/horus/res"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email       string             `json:"email" bson:"email" validate:"required,email"`
	FirstName   string             `json:"firstName" bson:"firstName" validate:"required"`
	LastName    string             `json:"lastName" bson:"lastName" validate:"required"`
	Password    string             `json:"password" bson:"password" validate:"required"`
	Verified    bool               `json:"verified" bson:"verified" validate:"required"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt" validate:"required"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt" validate:"required"`
	LastLoginAt time.Time          `json:"lastLoginAt,omitempty" bson:"lastLoginAt,omitempty"`
}

func (m *User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type userIDCtxKey struct{}

type UserHandler struct {
	Logger    *logger.Logger
	Validator *validator.Validator
	DB        *mongo.Database
}

func (h *UserHandler) Routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/", h.Find) // GET /user
	mux.Route("/{id}", func(mux chi.Router) {
		mux.Use(middlewares.ValidateObjectID("id"))
		mux.Use(UserCtx)
		mux.Get("/", h.FindById) // GET /user/:id
	})
	return mux
}

func (h *UserHandler) Find(rw http.ResponseWriter, r *http.Request) {
	// fetch
	collection := h.DB.Collection("users")
	cursor, err := collection.Find(r.Context(), bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusOK)
			render.Render(rw, r, &res.JSON{"users": []interface{}{}})
			return
		}
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// decode
	var users []*User
	if err = cursor.All(r.Context(), &users); err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// compile renders
	renders := []render.Renderer{}
	for _, user := range users {
		renders = append(renders, user)
	}

	// render
	resp := &res.JSON{"users": renders}
	render.Status(r, http.StatusOK)
	if err := render.Render(rw, r, resp); err != nil {
		render.Render(rw, r, res.ErrRender(err))
		return
	}
}

func (h *UserHandler) FindById(rw http.ResponseWriter, r *http.Request) {
	// serialize id
	id := r.Context().Value(userIDCtxKey{}).(string)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// fetch
	collection := h.DB.Collection("users")
	result := collection.FindOne(r.Context(), bson.M{"_id": oid})
	err = result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Render(rw, r, res.ErrNotFound(errors.New("no user found")))
			return
		}
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// decode
	var user *User
	err = result.Decode(&user)
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// render
	resp := &res.JSON{"user": user}
	render.Status(r, http.StatusOK)
	if err := render.Render(rw, r, resp); err != nil {
		render.Render(rw, r, res.ErrRender(err))
		return
	}
}

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), userIDCtxKey{}, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.Clone(ctx))
	})
}

func NewUserHandler(logger *logger.Logger, validator *validator.Validator, db *mongo.Database) *UserHandler {
	return &UserHandler{
		Logger:    logger,
		Validator: validator,
		DB:        db,
	}
}
