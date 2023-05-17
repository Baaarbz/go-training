package finadbyid

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	"barbz.dev/marketplace/pkg/valueobject"
)

type FindAdById struct {
	ads AdRepository
}

func NewFindAdById(ads AdRepository) FindAdById {
	return FindAdById{
		ads: ads,
	}
}

type GetAdByIdResponse struct {
	Id          string
	Title       string
	Description string
	Price       float32
	Date        string
}

func (service FindAdById) Execute(id string) (response GetAdByIdResponse, err error) {
	adId, err := valueobject.NewId(id)
	ad, err := service.ads.FindAdById(adId)
	response = mapToResponse(ad)
	return
}

func mapToResponse(ad Ad) GetAdByIdResponse {
	return GetAdByIdResponse{
		Id:          ad.GetId().String(),
		Title:       ad.Title.String(),
		Description: ad.Description.String(),
		Price:       ad.Price.Number(),
		Date:        ad.GetDate().String(),
	}
}
