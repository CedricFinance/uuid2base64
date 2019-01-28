package main

import "testing"

func TestConvertToBase64(t *testing.T) {
	type args struct {
		uuidStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{ "valid uuid", args{"b0ba9400-e618-11e6-a000-0031f9823668"}, "sLqUAOYYEeagAAAx+YI2aA==", false },
		{ "invalid uuid", args{"invalid"}, "", true },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToBase64(tt.args.uuidStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertToBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertToBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
