package repository

import (
	"barbz.dev/marketplace/internal/pkg/domain"
	"errors"
	"github.com/google/uuid"
)

var ads = make([]domain.Ad, 0)

func Save(ad domain.Ad) domain.Ad {
	var id, _ = uuid.NewUUID()
	ad.SetID(id.String())

	ads = append(ads, ad)
	return ad
}

func FindById(id string) (domain.Ad, error) {
	for _, ad := range ads {
		if ad.GetId() == id {
			return ad, nil
		}
	}
	return domain.Ad{}, errors.New("not found Ad")
}
