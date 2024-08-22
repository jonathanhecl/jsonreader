package jsonreader

import (
	"reflect"
	"testing"
)

func TestLoadFileJSON(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    JSONStruct
		wantErr bool
	}{
		{
			name: "TestLoadFileJSON",
			args: args{
				filename: "example.json",
			},
			want: JSONStruct{
				Headers: []string{"id", "father", "mother", "children"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFileJSON(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFileJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadFileJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
