package ad

import "barbz.dev/marketplace/pkg/valueobject"

type AdRepository interface {
	SaveAd(ad Ad) Ad
	FindAdById(id valueobject.AdId) (Ad, error)
	FindAllAds() (adResponse []Ad)
}
