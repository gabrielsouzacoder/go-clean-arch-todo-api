package entity

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestStringToID(t *testing.T) {
	type args struct {
		s string
	}
	var (
		parse, _ = uuid.Parse("e3c584de-19d2-4de8-b6a5-ac8766589003")
		tests    = []struct {
			name    string
			args    args
			want    ID
			wantErr bool
		}{
			{
				name: "Should convert string to ID",
				args: args{
					s: "e3c584de-19d2-4de8-b6a5-ac8766589003",
				},
				want:    parse,
				wantErr: false,
			},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToID(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
