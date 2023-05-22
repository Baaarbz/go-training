package ad

import (
	"barbz.dev/marketplace/internal/infrastructure/server/configuration"
	"barbz.dev/marketplace/internal/pkg/application/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveAdHandler struct {
	saveAd ad.SaveAd
}

func NewSaveAdHandler(dependencies *configuration.AdConfiguration) *SaveAdHandler {
	return &SaveAdHandler{
		saveAd: dependencies.SaveAdService,
	}
}

type JSONSaveAdResponse struct {
	Id string `json:"id"`
}

func (h SaveAdHandler) SaveAd() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ad.SaveAdRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if savedAd, err := h.saveAd.Execute(ctx, req); err != nil {
			switch {
			case errors.Is(err, ErrAdIdBadFormat) ||
				errors.Is(err, ErrTitleBadFormat) ||
				errors.Is(err, ErrDescriptionBadFormat) ||
				errors.Is(err, ErrPriceBadFormat):
				ctx.String(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			ctx.JSON(http.StatusCreated, h.mapSaveAdToJsonResponse(savedAd))
		}
	}
}

func (SaveAdHandler) mapSaveAdToJsonResponse(response ad.SaveAdResponse) JSONSaveAdResponse {
	return JSONSaveAdResponse{
		Id: response.Id,
	}
}
