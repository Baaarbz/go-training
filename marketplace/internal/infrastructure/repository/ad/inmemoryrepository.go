package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"errors"
	"github.com/google/uuid"
)

type InMemoryRepository struct {
	ads []Ad
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		ads: make([]Ad, 0),
	}
}

func (repository *InMemoryRepository) SaveAd(ad Ad) (Ad, error) {
	var id, _ = uuid.NewUUID()
	var adId, _ = NewId(id.String())
	ad.SetId(adId)

	repository.ads = append(repository.ads, ad)
	return ad, nil
}

func (repository *InMemoryRepository) FindAdById(id AdId) (Ad, error) {
	for _, ad := range repository.ads {
		if ad.GetId() == id {
			return ad, nil
		}
	}
	return Ad{}, errors.New("not found Ad")
}

func (repository *InMemoryRepository) FindAllAds() (adResponse []Ad, err error) {
	if len(repository.ads) < 5 {
		adResponse = repository.ads
	} else {
		adResponse = repository.ads[:5]
	}
	return
}
