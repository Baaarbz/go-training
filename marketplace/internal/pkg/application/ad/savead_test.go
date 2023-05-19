package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	"barbz.dev/marketplace/internal/pkg/domain/ad/mocks"
	. "barbz.dev/marketplace/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSaveAd_Execute_Success(t *testing.T) {
	ads := mocks.NewAdRepository(t)
	service := NewSaveAd(ads)

	requestAd := SaveAdRequest{
		Title:       "Test save ad success",
		Description: "Description to save successfully an ad",
		Price:       99.99,
	}
	savedAd := NewAd(Title(requestAd.Title), Description(requestAd.Description), Price(requestAd.Price))
	savedAd.SetId("test-id")
	ads.EXPECT().SaveAd(mock.AnythingOfType("Ad")).Return(savedAd, nil)

	gotAd, err := service.Execute(requestAd)

	assert.Nil(t, err)
	assert.Equal(t, SaveAdResponse{savedAd.GetId().String()}, gotAd)
}

func TestSaveAd_Execute_FailValidation(t *testing.T) {
	ads := mocks.NewAdRepository(t)
	type args struct {
		request SaveAdRequest
	}
	tests := []struct {
		name    string
		args    args
		want    SaveAdResponse
		wantErr bool
	}{
		{"wrong Title format", args{request: SaveAdRequest{Title: ":(", Description: "Description to pass the validation", Price: 0}}, SaveAdResponse{}, true},
		{"wrong Title format", args{request: SaveAdRequest{Title: "Title to pass validation", Description: ":(", Price: 0}}, SaveAdResponse{}, true},
		{"wrong Title format", args{request: SaveAdRequest{Title: "Title to pass validation", Description: "Description to pass the validation", Price: -20}}, SaveAdResponse{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := SaveAd{ads: ads}
			gotSavedAd, err := service.Execute(tt.args.request)
			assert.Equal(t, SaveAdResponse{}, gotSavedAd)
			assert.Error(t, err)
		})
	}
}
