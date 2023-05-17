package valueobject

import "testing"

func TestNewId(t *testing.T) {
	id, err := NewId("574cc928-f4bd-11ed-ad0e-8a6a68a798d6")

	if err != nil {
		t.Errorf("id %v has wrong format: %v", id, err)
	}
}
func TestNewIdWrongFormat(t *testing.T) {
	id, err := NewId("wrong-id-format")

	if err == nil {
		t.Errorf("id %v has invalid format (should throw error)", id)
	}
}
