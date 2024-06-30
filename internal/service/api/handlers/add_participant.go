package handlers

import (
	"errors"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/black-pepper-team/community-indexer/internal/service/api/requests"
	"github.com/black-pepper-team/community-indexer/internal/service/api/responses"
	"github.com/black-pepper-team/community-indexer/internal/service/core"
)

func AddCommunityParticipant(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewAddCommunityParticipant(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	mintRequest, err := Core(r).AddCommunityParticipant(
		req.PrivateKey, req.ParticipantAddress, req.ContractAddress,
	)
	switch {
	case errors.Is(err, core.ErrContractNotFound):
		ape.RenderErr(w, problems.NotFound())
		Log(r).WithError(err).
			Debug("Contract not found")
		return
	case err != nil:
		Log(r).WithError(err).
			Error("Failed add community participant")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewRegisterInCommunity(mintRequest))
}
