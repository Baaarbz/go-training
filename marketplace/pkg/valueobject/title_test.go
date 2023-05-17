package valueobject

import "testing"

func TestNewTitle(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name        string
		args        args
		wantAdTitle Title
		wantErr     bool
	}{
		{"empty title should error", args{""}, "", true},
		{"3 letters title should error", args{"err"}, "", true},
		{"6 or more letters should success", args{"Ad Title"}, "Ad Title", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdTitle, err := NewTitle(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAdTitle != tt.wantAdTitle {
				t.Errorf("NewTitle() gotAdTitle = %v, want %v", gotAdTitle, tt.wantAdTitle)
			}
		})
	}
}
