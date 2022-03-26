package idgen

import (
	"testing"
	"time"
)

func TestGenerateID(t *testing.T) {
	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		{name: "case1", want: time.Now().Unix(), wantErr: false},
	}
	gen := NewRandomIDGen()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gen.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got/100 != tt.want/100 {
				t.Errorf("GenerateID() / 100 got = %v, want %v", got, tt.want)
			}
		})
	}
}
