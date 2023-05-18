package ad

import "barbz.dev/marketplace/pkg/valueobject"

//go:generate mockery --name=AdRepository --filename mock_adrepository.go
type AdRepository interface {
	SaveAd(ad Ad) Ad
	FindAdById(id valueobject.AdId) (Ad, error)
	FindAllAds() (adResponse []Ad)
}
