package organizations

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/horus/logger"
	"github.com/knuls/horus/middlewares"
	"github.com/knuls/horus/res"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type organizationIDCtxKey struct{}

type handler struct {
	logger *logger.Logger
	dao    *dao
}

func (h *handler) Routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/", h.Find)    // GET /organization
	mux.Post("/", h.Create) // POST /organization
	mux.Route("/{id}", func(mux chi.Router) {
		mux.Use(middlewares.ValidateObjectID("id"))
		mux.Use(organizationCtx)
		mux.Get("/", h.FindById) // GET /organization/:id
	})
	return mux
}

func (h *handler) Find(rw http.ResponseWriter, r *http.Request) {
	orgs, err := h.dao.Find(r.Context(), bson.D{})
	if err != nil {
		h.logger.Error("failed to find organizations", "error", err)
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	renders := []render.Renderer{}
	for _, org := range orgs {
		renders = append(renders, org)
	}
	render.Status(r, http.StatusOK)
	if err = render.Render(rw, r, &res.JSON{"organizations": renders}); err != nil {
		h.logger.Error("failed to render", "error", err)
		render.Render(rw, r, res.ErrRender(err))
		return
	}
}

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {
	org := NewOrganization()
	defer r.Body.Close()
	if err := org.FromJSON(r.Body); err != nil {
		h.logger.Error("failed to decode request body", "error", err)
		render.Render(rw, r, res.ErrDecode(err))
		return
	}
	id, err := h.dao.Create(r.Context(), org)
	if err != nil {
		h.logger.Error("failed to create organization", "error", err)
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	if err = render.Render(rw, r, &res.JSON{"id": id}); err != nil {
		h.logger.Error("failed to render", "error", err)
	}
}

func organizationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), organizationIDCtxKey{}, chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.Clone(ctx))
	})
}

func (h *handler) FindById(rw http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(organizationIDCtxKey{}).(string)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		h.logger.Error("failed to convert hex to object id", "error", err)
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	org, err := h.dao.FindOne(r.Context(), bson.D{{Key: "_id", Value: oid}})
	if err != nil {
		h.logger.Error("failed to find organization", "error", err)
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	if err := render.Render(rw, r, &res.JSON{"organization": org}); err != nil {
		h.logger.Error("failed to render", "error", err)
		render.Render(rw, r, res.ErrRender(err))
		return
	}
}

func NewHandler(logger *logger.Logger, dao *dao) *handler {
	return &handler{
		logger: logger,
		dao:    dao,
	}
}
