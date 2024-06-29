package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/black-pepper-team/community-indexer/internal/service/api/requests"
	"github.com/black-pepper-team/community-indexer/internal/service/api/responses"
)

func CreateCommunity(w http.ResponseWriter, r *http.Request) {
	if MockAPI(r) {
		ape.Render(w, responses.MockedCreateCommunity())
		return
	}

	req, err := requests.NewCreateCommunity(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	newCommunity, err := Core(r).CreateCommunity(req.CollectionName, req.CollectionSymbol)
	if err != nil {
		Log(r).WithError(err).
			Error("Failed create community")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewCreateCommunity(newCommunity))
}
