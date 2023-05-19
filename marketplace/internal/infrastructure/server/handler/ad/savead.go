package ad

import (
	"barbz.dev/marketplace/internal/pkg/application/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveAd(service ad.SaveAd) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ad.SaveAdRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if savedAd, err := service.Execute(ctx, req); err != nil {
			switch {
			case errors.Is(err, ErrAdIdBadFormat):
			case errors.Is(err, ErrTitleBadFormat):
			case errors.Is(err, ErrDateBadFormat):
			case errors.Is(err, ErrDescriptionBadFormat):
			case errors.Is(err, ErrPriceBadFormat):
				ctx.JSON(http.StatusBadRequest, err)
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			ctx.JSON(http.StatusCreated, savedAd)
		}
	}
}
