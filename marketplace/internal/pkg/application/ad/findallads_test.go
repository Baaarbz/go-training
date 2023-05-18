package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	"barbz.dev/marketplace/internal/pkg/domain/ad/mocks"
	"barbz.dev/marketplace/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAllAds_Execute_GetListOfAds(t *testing.T) {
	ads := mocks.NewAdRepository(t)
	service := NewFindAllAds(ads)

	anId, _ := valueobject.NewId("574cc928-f4bd-11ed-ad0e-8a6a68a798d6")
	ad := NewAd("Simple title", "Simple ad description for testing", 20)
	ad.SetId(anId)
	ads.EXPECT().FindAllAds().Return([]Ad{ad}, nil)

	gotAds, err := service.Execute()

	assert.True(t, len(gotAds) == 1)
	assert.Nil(t, err)
	assert.Equal(t, []GetAdsResponse{{anId.String()}}, gotAds)
}
