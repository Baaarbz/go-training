package ad

type Repository interface {
	SaveAd(ad Ad) Ad
	FindAdById(id string) (Ad, error)
	FindAllAds() (adResponse []Ad)
}
