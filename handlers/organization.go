package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/bennu/models"
	"github.com/knuls/horus/logger"
	"github.com/knuls/horus/middlewares"
	"github.com/knuls/horus/res"
	"github.com/knuls/horus/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type organizationIDCtxKey struct{}

type OrganizationHandler struct {
	Logger    *logger.Logger
	Validator *validator.Validator
	DB        *mongo.Database
}

func (h *OrganizationHandler) Routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/", h.Find)    // GET /organization
	mux.Post("/", h.Create) // POST /organization
	mux.Route("/{id}", func(mux chi.Router) {
		mux.Use(middlewares.ValidateObjectID("id"))
		mux.Use(OrganizationCtx)
		mux.Get("/", h.FindById) // GET /organization/:id
	})
	return mux
}

func (h *OrganizationHandler) Find(rw http.ResponseWriter, r *http.Request) {
	// fetch
	collection := h.DB.Collection("organizations")
	cursor, err := collection.Find(r.Context(), bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Status(r, http.StatusOK)
			render.Render(rw, r, &res.JSON{"organizations": []interface{}{}})
			return
		}
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// decode
	var orgs []*models.Organization
	if err = cursor.All(r.Context(), &orgs); err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// compile renders
	renders := []render.Renderer{}
	for _, org := range orgs {
		renders = append(renders, org)
	}

	// render
	resp := &res.JSON{"organizations": renders}
	render.Status(r, http.StatusOK)
	if err := render.Render(rw, r, resp); err != nil {
		render.Render(rw, r, res.ErrRender(err))
		return
	}
}

func (h *OrganizationHandler) Create(rw http.ResponseWriter, r *http.Request) {
	// decode
	var org *models.Organization
	err := json.NewDecoder(r.Body).Decode(&org)
	defer r.Body.Close()
	if err == io.EOF {
		render.Render(rw, r, res.ErrDecode(err))
		return
	}
	if err != nil {
		render.Render(rw, r, res.ErrDecode(err))
		return
	}

	// validate
	if err := h.Validator.ValidateStruct(org); err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// orgs
	collection := h.DB.Collection("organizations")

	// ensure email does not exist
	count, err := collection.CountDocuments(r.Context(), bson.M{"name": org.Name})
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	if count > 0 {
		render.Render(rw, r, res.ErrBadRequest(errors.New("name already exists")))
		return
	}

	// insert
	result, err := collection.InsertOne(r.Context(), org)
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// render
	render.Status(r, http.StatusCreated)
	render.Respond(rw, r, &res.JSON{"id": result.InsertedID})
}

func (h *OrganizationHandler) FindById(rw http.ResponseWriter, r *http.Request) {
	// serialize id
	id := r.Context().Value(organizationIDCtxKey{}).(string)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// fetch
	collection := h.DB.Collection("organizations")
	result := collection.FindOne(r.Context(), bson.M{"_id": oid})
	err = result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			render.Render(rw, r, res.ErrNotFound(errors.New("no organization found")))
			return
		}
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// decode
	var org *models.Organization
	err = result.Decode(&org)
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}

	// render
	resp := &res.JSON{"organization": org}
	render.Status(r, http.StatusOK)
	if err := render.Render(rw, r, resp); err != nil {
		render.Render(rw, r, res.ErrRender(err))
		return
	}
}

func OrganizationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), organizationIDCtxKey{}, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.Clone(ctx))
	})
}

func NewOrganizationHandler(logger *logger.Logger, validator *validator.Validator, db *mongo.Database) *OrganizationHandler {
	return &OrganizationHandler{
		Logger:    logger,
		Validator: validator,
		DB:        db,
	}
}
