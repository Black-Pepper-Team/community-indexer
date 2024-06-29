package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/black-pepper-team/community-indexer/internal/service/api/responses"
)

func CommunitiesList(w http.ResponseWriter, r *http.Request) {
	if MockAPI(r) {
		ape.Render(w, responses.MockedCommunitiesList())
		return
	}

	communitiesList, err := Core(r).GetCommunitiesList()
	if err != nil {
		Log(r).WithError(err).
			Error("Failed get communities list")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewCommunitiesList(communitiesList))
}