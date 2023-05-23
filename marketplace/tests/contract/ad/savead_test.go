package ad

import (
	"barbz.dev/marketplace/internal/pkg/application/ad"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveAd(t *testing.T) {
	serviceExpectedResponse := ad.SaveAdDtoResponse{Id: "test-id"}
	expectedSaveAdJsonResponse := `{"id": "test-id"}`
	body := []byte(`{"title": "iPhone", "description": "Description of iPhone test", "price": 1000}`)
	saveAdMock.EXPECT().Execute(mock.Anything, mock.AnythingOfType("ad.SaveAdDtoRequest")).Return(serviceExpectedResponse, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/ads", bytes.NewReader(body))
	srv.Engine.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.JSONEq(t, expectedSaveAdJsonResponse, w.Body.String())
}
