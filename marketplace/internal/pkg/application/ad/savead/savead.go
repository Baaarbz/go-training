package savead

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
)

type SaveAd struct {
	ads AdRepository
}

func NewSaveAd(ads AdRepository) SaveAd {
	return SaveAd{
		ads: ads,
	}
}

type SaveAdResponse struct {
	Id string
}

type SaveAdRequest struct {
	Title       string  `json:"title" required:"true"`
	Description string  `json:"description" required:"true"`
	Price       float32 `json:"price" required:"true"`
}

func (service SaveAd) Execute(request SaveAdRequest) (SaveAdResponse, error) {
	if title, description, price, err := getFieldsAds(request); err != nil {
		return SaveAdResponse{}, err
	} else {
		ad := NewAd(title, description, price)
		ad = service.ads.SaveAd(ad)
		return SaveAdResponse{Id: ad.GetId().String()}, nil
	}
}

func getFieldsAds(request SaveAdRequest) (title Title, description Description, price Price, err error) {
	title, err = NewTitle(request.Title)
	description, err = NewDescription(request.Description)
	price, err = NewPrice(request.Price)
	return
}
