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
				Headers: []string{"id", "father", "mother", "children", "extra"},
				Rows: map[int]map[string]string{
					0: {"id": "1", "father": "Mark", "mother": "Charlotte", "children": "2"},
					1: {"id": "2", "father": "John", "mother": "Ann", "children": "3", "extra": "1"},
					2: {"id": "3", "father": "Bob", "mother": "Monika", "children": "2"},
				},
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

			if !reflect.DeepEqual(got.Headers, tt.want.Headers) {
				t.Errorf("LoadFileJSON() got.Headers = %v, want %v", got.Headers, tt.want.Headers)
			}

			for i := range got.Rows {
				if !reflect.DeepEqual(got.Rows[i], tt.want.Rows[i]) {
					t.Errorf("LoadFileJSON() got.Rows[%v] = %v, want %v", i, got.Rows[i], tt.want.Rows[i])
				}
			}
		})
	}
}
