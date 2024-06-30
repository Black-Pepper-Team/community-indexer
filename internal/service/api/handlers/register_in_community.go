package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/black-pepper-team/community-indexer/internal/service/api/requests"
	"github.com/black-pepper-team/community-indexer/internal/service/api/responses"
)

func RegisterInCommunity(w http.ResponseWriter, r *http.Request) {
	if MockAPI(r) {
		ape.Render(w, responses.MockedCreateCommunity())
		return
	}

	req, err := requests.NewRegisterInCommunity(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	registerRequest, err := Core(r).RegisterInCommunity(
		req.BJJPublicKey,
		req.NFTOwner,
		req.ContractId,
		req.NFTID,
	)
	if err != nil {
		Log(r).WithError(err).
			Error("Failed register in community")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewRegisterInCommunity(registerRequest))
}
