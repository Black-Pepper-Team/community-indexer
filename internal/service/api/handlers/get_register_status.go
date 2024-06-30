package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/black-pepper-team/community-indexer/internal/service/api/requests"
	"github.com/black-pepper-team/community-indexer/internal/service/api/responses"
)

func GetRegisterStatus(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewGetRegister(r)
	if err != nil {
		Log(r).WithField("reason", err).Debug("Bad request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	registryRecord, err := Core(r).GetRegister(req.RegisterID)
	switch {
	case registryRecord == nil:
		ape.RenderErr(w, problems.NotFound())
		Log(r).WithError(err).
			Debug("Community not found")
		return
	case err != nil:
		Log(r).WithError(err).
			Error("Failed to get register status")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewGetRegisterStatus(registryRecord))
}
