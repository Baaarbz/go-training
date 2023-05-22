package ad

import (
	"barbz.dev/marketplace/internal/pkg/application/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONSaveAdResponse struct {
	Id string `json:"id"`
}

func SaveAd(service ad.SaveAd) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ad.SaveAdRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if savedAd, err := service.Execute(ctx, req); err != nil {
			switch {
			case errors.Is(err, ErrAdIdBadFormat) ||
				errors.Is(err, ErrTitleBadFormat) ||
				errors.Is(err, ErrDateBadFormat) ||
				errors.Is(err, ErrDescriptionBadFormat) ||
				errors.Is(err, ErrPriceBadFormat):
				ctx.String(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			ctx.JSON(http.StatusCreated, mapSaveAdToJsonResponse(savedAd))
		}
	}
}

func mapSaveAdToJsonResponse(response ad.SaveAdResponse) JSONSaveAdResponse {
	return JSONSaveAdResponse{
		Id: response.Id,
	}
}
