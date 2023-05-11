package main

import (
	"barbz.dev/marketplace/internal/infrastructure/repository"
	"barbz.dev/marketplace/internal/pkg/domain"
	"log"
)

func main() {
	ad := domain.NewAd("Sample Ad", "This is the description to test the app", 29.99)
	ad = repository.SaveAd(ad)
	addMocksToRepository()

	printFindById(ad.GetId())
	printFindFiveAds()
}

func addMocksToRepository() {
	ad := domain.NewAd("Sample Ad", "This is the description to test the app", 29.99)
	repository.SaveAd(ad)
	ad = domain.NewAd("BMW 320D", "This is the description to test the app", 1800.99)
	repository.SaveAd(ad)
	ad = domain.NewAd("TV 45' Sony", "This is the description to test the app", 399.99)
	repository.SaveAd(ad)
	ad = domain.NewAd("Sportiva Rock Climbing Foot", "This is the description to test the app", 70.00)
	repository.SaveAd(ad)
	ad = domain.NewAd("Macbook Pro 16'", "This is the description to test the app", 1500.00)
	repository.SaveAd(ad)
	ad = domain.NewAd("Rolex limited edition", "This is the description to test the app", 29999.99)
	repository.SaveAd(ad)
}

func printFindById(id string) {
	if adFound, err := repository.FindAdById(id); err != nil {
		log.Fatalf("exception: %v\n", err)
	} else {
		log.Println(adFound)
	}
}

func printFindFiveAds() {
	ads := repository.FindFiveAds()
	log.Println("Size of the array:", len(ads))
	log.Println(ads)
}
