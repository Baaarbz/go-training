package bootstrap

import (
	. "barbz.dev/marketplace/internal/infrastructure/repository/ad"
	. "barbz.dev/marketplace/internal/pkg/application/ad"
	"log"
)

func Run() {
	saveAd, findAllAds, findAdById := initAdsDependencies()
	runMocks(saveAd, findAllAds, findAdById)
}

func initAdsDependencies() (save SaveAd, findAll FindAllAds, findById FindAdById) {
	ads := NewInMemoryRepository()
	save = NewSaveAd(ads)
	findAll = NewFindAllAds(ads)
	findById = NewFindAdById(ads)
	return
}

func runMocks(saveAd SaveAd, findAllAds FindAllAds, findAdById FindAdById) {
	ad := SaveAdRequest{
		Title:       "Sample Ad",
		Description: "This is the description to test the app",
		Price:       29.99,
	}
	response, _ := saveAd.Execute(ad)
	addMocksToRepository(saveAd)

	if adFound, err := findAdById.Execute(response.Id); err != nil {
		log.Fatalf("exception: %v\n", err)
	} else {
		log.Println(adFound)
	}

	allAds := findAllAds.Execute()
	log.Println("Size of the array:", len(allAds))
	log.Println(allAds)
}

func addMocksToRepository(saveAd SaveAd) {
	ad := SaveAdRequest{Title: "Sample Ad", Description: "This is the description to test the app", Price: 29.99}
	saveAd.Execute(ad)
	ad = SaveAdRequest{Title: "BMW 320D", Description: "This is the description to test the app", Price: 1800.99}
	saveAd.Execute(ad)
	ad = SaveAdRequest{Title: "TV 45' Sony", Description: "This is the description to test the app", Price: 399.99}
	saveAd.Execute(ad)
	ad = SaveAdRequest{Title: "Sportiva Rock Climbing Foot", Description: "This is the description to test the app", Price: 70.00}
	saveAd.Execute(ad)
	ad = SaveAdRequest{Title: "Macbook Pro 16'", Description: "This is the description to test the app", Price: 1500.00}
	saveAd.Execute(ad)
	ad = SaveAdRequest{Title: "Rolex limited edition", Description: "This is the description to test the app", Price: 29999.99}
	saveAd.Execute(ad)
}
