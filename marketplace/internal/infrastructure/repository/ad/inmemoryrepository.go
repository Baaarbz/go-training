package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"errors"
	"github.com/google/uuid"
)

var ads = make([]Ad, 0)

func SaveAd(ad Ad) Ad {
	var id, _ = uuid.NewUUID()
	var adId, _ = NewId(id.String())
	ad.SetID(adId)

	ads = append(ads, ad)
	return ad
}

func FindAdById(id AdId) (Ad, error) {
	for _, ad := range ads {
		if ad.GetId() == id {
			return ad, nil
		}
	}
	return Ad{}, errors.New("not found Ad")
}

func FindAllAds() (adResponse []Ad) {
	if len(ads) < 5 {
		adResponse = ads
	} else {
		adResponse = ads[:5]
	}
	return
}
