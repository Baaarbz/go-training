package main

import (
	"barbz.dev/marketplace/internal/infrastructure/repository"
	"barbz.dev/marketplace/internal/pkg/domain"
	"fmt"
)

func main() {
	ad := domain.NewAd("Sample Ad", "This is the description to test the app", 29.99)
	ad = repository.Save(ad)

	printFindById(ad.GetId())
}

func printFindById(id string) {
	if adFound, err := repository.FindById(id); err != nil {
		fmt.Printf("exception: %v\n", err)
	} else {
		fmt.Println(adFound)
	}
}
