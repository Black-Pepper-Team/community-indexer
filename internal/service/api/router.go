package api

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"github.com/black-pepper-team/community-indexer/internal/service/api/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxCore(s.core),
			handlers.CtxMockAPI(s.mockAPI),
		),
	)
	r.Route("/integrations/community-indexer/v1", func(r chi.Router) {
		r.Route("/community", func(r chi.Router) {
			r.Get("/{community-id}", handlers.GetCommunity)
			r.Post("/", handlers.CreateCommunity)
			r.Get("/list", handlers.CommunityList)
			r.Post("/import", handlers.ImportCommunity)
			r.Post("/register", handlers.RegisterInCommunity)
			r.Get("/register/{register-id}", handlers.GetRegisterStatus)
		})
	})

	return r
}
