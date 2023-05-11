package repository

import (
	. "barbz.dev/marketplace/internal/pkg/domain"
	"errors"
	"github.com/google/uuid"
)

var ads = make([]Ad, 0)

func SaveAd(ad Ad) Ad {
	var id, _ = uuid.NewUUID()
	ad.SetID(id.String())

	ads = append(ads, ad)
	return ad
}

func FindAdById(id string) (Ad, error) {
	for _, ad := range ads {
		if ad.GetId() == id {
			return ad, nil
		}
	}
	return Ad{}, errors.New("not found Ad")
}

func FindFiveAds() []Ad {
	return ads[:5]
}
