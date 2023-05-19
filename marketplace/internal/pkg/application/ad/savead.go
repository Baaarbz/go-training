//go:generate mockery --name=SaveAd --filename mock_savead.go
package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"context"
)

type SaveAd interface {
	Execute(ctx context.Context, request SaveAdRequest) (SaveAdResponse, error)
}

type saveAd struct {
	ads AdRepository
}

func NewSaveAd(ads AdRepository) saveAd {
	return saveAd{
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

func (service saveAd) Execute(ctx context.Context, request SaveAdRequest) (SaveAdResponse, error) {
	if title, description, price, err := service.getFieldsAds(request); err != nil {
		return SaveAdResponse{}, err
	} else {
		ad := NewAd(title, description, price)
		ad, err = service.ads.SaveAd(ctx, ad)
		return SaveAdResponse{Id: ad.GetId().String()}, err
	}
}

func (saveAd) getFieldsAds(request SaveAdRequest) (title Title, description Description, price Price, err error) {
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
