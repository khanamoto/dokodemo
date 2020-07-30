package web

import (
	"testing"
)

func Test_validateAll(t *testing.T) {
	type args struct {
		dataSet interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAll(tt.args.dataSet); (err != nil) != tt.wantErr {
				t.Errorf("validateAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateUserBase(t *testing.T) {
	type args struct {
		dataSet interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUserBase(tt.args.dataSet); (err != nil) != tt.wantErr {
				t.Errorf("validateUserBase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateUserName(t *testing.T) {
	type args struct {
		dataSet interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUserName(tt.args.dataSet); (err != nil) != tt.wantErr {
				t.Errorf("validateUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
