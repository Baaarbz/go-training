package ad

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
	Title       string
	Description string
	Price       float32
}

func (service SaveAd) Execute(request SaveAdRequest) (SaveAdResponse, error) {
	if title, description, price, err := service.getFieldsAds(request); err != nil {
		return SaveAdResponse{}, err
	} else {
		ad := NewAd(title, description, price)
		ad = service.ads.SaveAd(ad)
		return SaveAdResponse{Id: ad.GetId().String()}, nil
	}
}

func (SaveAd) getFieldsAds(request SaveAdRequest) (title Title, description Description, price Price, err error) {
	title, err = NewTitle(request.Title)
	if err != nil {
		return "", "", 0, err
	}
	description, err = NewDescription(request.Description)
	if err != nil {
		return "", "", 0, err
	}
	price, err = NewPrice(request.Price)
	return
}
