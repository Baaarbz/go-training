package ad

import . "barbz.dev/marketplace/internal/pkg/domain/ad"

type FindAllAds struct {
	ads AdRepository
}

func NewFindAllAds(ads AdRepository) FindAllAds {
	return FindAllAds{
		ads: ads,
	}
}

type GetAdsResponse struct {
	Id string
}

func (service FindAllAds) Execute() ([]GetAdsResponse, error) {
	ads, err := service.ads.FindAllAds()
	return service.mapToResponse(ads), err
}

func (FindAllAds) mapToResponse(ads []Ad) []GetAdsResponse {
	adsResponse := make([]GetAdsResponse, 0)
	for _, ad := range ads {
		adsResponse = append(adsResponse, GetAdsResponse{Id: ad.GetId().String()})
	}
	return adsResponse
}
