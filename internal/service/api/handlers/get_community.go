package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/black-pepper-team/community-indexer/internal/service/api/requests"
	"github.com/black-pepper-team/community-indexer/internal/service/api/responses"
)

func GetCommunity(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetCommunity(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if MockAPI(r) {
		ape.Render(w, responses.MockedGetCommunity(req))
		return
	}

	community, err := Core(r).GetCommunityById(req.CommunityID)
	switch {
	case community == nil:
		ape.RenderErr(w, problems.NotFound())
		Log(r).WithError(err).
			Debug("Community not found")
		return
	case err != nil:
		Log(r).WithError(err).
			Error("Failed tp get community")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewGetCommunity(community))
}
