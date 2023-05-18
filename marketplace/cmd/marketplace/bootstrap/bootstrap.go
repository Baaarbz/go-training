package bootstrap

import (
	. "barbz.dev/marketplace/internal/infrastructure/repository/ad"
	. "barbz.dev/marketplace/internal/pkg/application/ad"
	"log"
)

func Run() {
	saveAd, findAllAds, findAdById := initAdsDependencies()
	log.Printf("%v\t%v\t%v", saveAd, findAllAds, findAdById)
}

func initAdsDependencies() (save SaveAd, findAll FindAllAds, findById FindAdById) {
	ads := NewInMemoryRepository()
	save = NewSaveAd(ads)
	findAll = NewFindAllAds(ads)
	findById = NewFindAdById(ads)
	return
}
