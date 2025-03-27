package endpoints

import (
	internalerrors "EmailNGo/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {
	campaigns, err := h.CampaignService.Repository.Get()

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}

		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, 200)
	render.JSON(w, r, campaigns)
}
