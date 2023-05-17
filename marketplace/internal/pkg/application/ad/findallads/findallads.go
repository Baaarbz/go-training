package findallads

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

func (service FindAllAds) Execute() []GetAdsResponse {
	ads := service.ads.FindAllAds()
	return mapToResponse(ads)
}

func mapToResponse(ads []Ad) []GetAdsResponse {
	adsResponse := make([]GetAdsResponse, 0)
	for _, ad := range ads {
		adsResponse = append(adsResponse, GetAdsResponse{Id: ad.GetId().String()})
	}
	return adsResponse
}
