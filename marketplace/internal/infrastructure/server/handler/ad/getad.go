package ad

import (
	"barbz.dev/marketplace/internal/pkg/application/ad"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllAds(service ad.FindAllAds) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ads, err := service.Execute(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		switch {
		case len(ads) == 0:
			ctx.Status(http.StatusNoContent)
			return
		default:
			ctx.JSON(http.StatusOK, ads)
			return
		}
	}
}

func FindAdById(service ad.FindAdById) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		adResponse, err := service.Execute(ctx, ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		switch {
		case len(adResponse.Id) == 0:
			ctx.Status(http.StatusNotFound)
			return
		default:
			ctx.JSON(http.StatusOK, adResponse)
			return
		}
	}
}
