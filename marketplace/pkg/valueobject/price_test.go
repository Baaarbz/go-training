package valueobject

import "testing"

func TestNewPrice(t *testing.T) {
	type args struct {
		value float32
	}
	tests := []struct {
		name        string
		args        args
		wantAdPrice Price
		wantErr     bool
	}{
		{"negative price should error", args{-10}, 0, true},
		{"price 0 should success", args{0}, 0, false},
		{"random positive price should success", args{120.20}, 120.20, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdPrice, err := NewPrice(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAdPrice != tt.wantAdPrice {
				t.Errorf("NewPrice() gotAdPrice = %v, want %v", gotAdPrice, tt.wantAdPrice)
			}
		})
	}
}
