//go:generate mockery --name=FindAllAds --filename mock_findallads.go
package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	"context"
)

const FindAllAdsBeanName = "ad.FindAllAds"

type FindAllAds interface {
	Execute(ctx context.Context) ([]GetAdsResponse, error)
}

type findAllAds struct {
	ads AdRepository
}

func NewFindAllAds(ads AdRepository) FindAllAds {
	return findAllAds{
		ads: ads,
	}
}

type GetAdsResponse struct {
	Id string
}

func (service findAllAds) Execute(ctx context.Context) ([]GetAdsResponse, error) {
	ads, err := service.ads.FindAllAds(ctx)
	return service.mapToResponse(ads), err
}

func (findAllAds) mapToResponse(ads []Ad) []GetAdsResponse {
	adsResponse := make([]GetAdsResponse, 0)
	for _, ad := range ads {
		adsResponse = append(adsResponse, GetAdsResponse{Id: ad.GetId().String()})
	}
	return adsResponse
}
