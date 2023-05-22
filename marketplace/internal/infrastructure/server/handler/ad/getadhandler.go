package ad

import (
	"barbz.dev/marketplace/internal/pkg/application/ad"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONFindAllAdsResponse struct {
	Id string `json:"id"`
}

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
			ctx.JSON(http.StatusOK, mapFindAllAdsToJSONResponse(ads))
			return
		}
	}
}

func mapFindAllAdsToJSONResponse(allAdsResponse []ad.GetAdsResponse) []JSONFindAllAdsResponse {
	jsonResponse := make([]JSONFindAllAdsResponse, 0)
	for _, response := range allAdsResponse {
		jsonResponse = append(jsonResponse, JSONFindAllAdsResponse{Id: response.Id})
	}
	return jsonResponse
}

type JSONFindAdByIdResponse struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Date        string  `json:"postedAt"`
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
			ctx.JSON(http.StatusOK, mapFindAdByIdToJSONResponse(adResponse))
			return
		}
	}
}

func mapFindAdByIdToJSONResponse(adByIdResponse ad.GetAdByIdResponse) JSONFindAdByIdResponse {
	return JSONFindAdByIdResponse{
		Id:          adByIdResponse.Id,
		Title:       adByIdResponse.Title,
		Description: adByIdResponse.Description,
		Price:       adByIdResponse.Price,
		Date:        adByIdResponse.Date,
	}
}
