package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/knuls/bennu/dao"
	"github.com/knuls/bennu/models"
	"github.com/knuls/horus/logger"
	"github.com/knuls/horus/middlewares"
	"github.com/knuls/horus/res"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type organizationIDCtxKey struct{}

type OrganizationHandler struct {
	logger     *logger.Logger
	daoFactory *dao.Factory
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
	orgs, err := h.daoFactory.GetOrganizationDao().Find(r.Context(), dao.Where{})
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	renders := []render.Renderer{}
	for _, org := range orgs {
		renders = append(renders, org)
	}
	render.Status(r, http.StatusOK)
	if err := render.Render(rw, r, &res.JSON{"organizations": renders}); err != nil {
		render.Render(rw, r, res.ErrRender(err))
		return
	}
}

func (h *OrganizationHandler) Create(rw http.ResponseWriter, r *http.Request) {
	var org *models.Organization
	defer r.Body.Close()
	if err := org.FromJSON(r.Body); err != nil {
		render.Render(rw, r, res.ErrDecode(err))
		return
	}
	id, err := h.daoFactory.GetOrganizationDao().Create(r.Context(), org)
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Respond(rw, r, &res.JSON{"id": id})
}

func (h *OrganizationHandler) FindById(rw http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(organizationIDCtxKey{}).(string)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	org, err := h.daoFactory.GetOrganizationDao().FindOne(r.Context(), dao.Where{{Key: "_id", Value: oid}})
	if err != nil {
		render.Render(rw, r, res.ErrBadRequest(err))
		return
	}
	render.Status(r, http.StatusOK)
	if err := render.Render(rw, r, &res.JSON{"organization": org}); err != nil {
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

func NewOrganizationHandler(logger *logger.Logger, factory *dao.Factory) *OrganizationHandler {
	return &OrganizationHandler{
		logger:     logger,
		daoFactory: factory,
	}
}
