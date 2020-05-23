package fetch

import (
	"archive/zip"
	"io"
	"testing"
	"time"
)

func TestSecUrl(t *testing.T) {
	type args struct {
		year    int
		quarter int
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{ "Quarter must be at least 1", args{2010, 0}, ErrQuarter },
		{ "Quarter must be less than 5", args{2010, 5}, ErrQuarter },
		{ "Year must be at least 2010", args{2008, 1}, ErrYear },
		{ "Year must be in the present", args{time.Now().Year() + 1, 1}, ErrYear },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SecUrl(tt.args.year, tt.args.quarter); err != tt.wantErr {
				t.Errorf("got = '%v', want '%v'", err, tt.wantErr)
			}
		})
	}
}

func TestUrl(t *testing.T) {
	type args struct {
		url string
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
			if err := Url(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Url() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestZip(t *testing.T) {
	type args struct {
		path string
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
			if err := Zip(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Zip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unzip(t *testing.T) {
	type args struct {
		r *zip.Reader
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
			if err := unzip(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("unzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_readZipFile(t *testing.T) {
	type args struct {
		file *zip.File
		fn   func(r io.Reader) error
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
			if err := readZipFile(tt.args.file, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("readZipFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
