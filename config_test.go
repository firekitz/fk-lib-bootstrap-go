package fkBootstrap

import (
	"fmt"
	"os"
	"testing"
)

type config struct {
	DO string `envconfig:"DO"`
}

func TestNoEnv(t *testing.T) {
	type args struct {
		cfg interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "No env",
			args:    args{cfg: &config{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadEnv(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("LoadEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("No env: %v\n", tt.args.cfg)
		})
	}
}

func TestLocalEnv(t *testing.T) {
	type args struct {
		cfg interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Local env",
			args:    args{cfg: &config{}},
			wantErr: false,
		},
	}
	os.Setenv("GO_ENV", "local")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadEnv(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("LoadEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("Local env: %v\n", tt.args.cfg)
		})
	}
}

func TestDeployEnv(t *testing.T) {
	type args struct {
		cfg interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Deploy env",
			args:    args{cfg: &config{}},
			wantErr: false,
		},
	}
	os.Setenv("GO_ENV", "dev")
	os.Setenv("DO", "deploy")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadEnv(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("LoadEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Printf("Deploy env: %v\n", tt.args.cfg)
		})
	}
}
