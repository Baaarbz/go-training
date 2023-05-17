package valueobject

import "testing"

func TestNewDescription(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name              string
		args              args
		wantAdDescription Description
		wantErr           bool
	}{
		{"empty description should error", args{"Short description error"}, "", true},
		{"3 words description should error", args{""}, "", true},
		{"4 words description should success", args{"4 words description should success"}, "4 words description should success", false},
		{"right format description should success", args{"right format description should success"}, "right format description should success", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdDescription, err := NewDescription(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAdDescription != tt.wantAdDescription {
				t.Errorf("NewDescription() gotAdDescription = %v, want %v", gotAdDescription, tt.wantAdDescription)
			}
		})
	}
}
