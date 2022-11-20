package bqclient

import (
	"bytes"
	"testing"
)

func TestQueryBasic(t *testing.T) {
	type args struct {
		projectID string
		sql       string
		location  string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := QueryBasic(w, tt.args.projectID, tt.args.sql, tt.args.location); (err != nil) != tt.wantErr {
				t.Errorf("QueryBasic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("QueryBasic() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_importCSVFromFile(t *testing.T) {
	type args struct {
		projectID string
		datasetID string
		tableID   string
		filename  string
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
			if err := ImportCSVFromFile(tt.args.projectID, tt.args.datasetID, tt.args.tableID, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("importCSVFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_importCSVExplicitSchema(t *testing.T) {
	type args struct {
		projectID string
		datasetID string
		tableID   string
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
			if err := ImportCSVExplicitSchema(tt.args.projectID, tt.args.datasetID, tt.args.tableID); (err != nil) != tt.wantErr {
				t.Errorf("importCSVExplicitSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
