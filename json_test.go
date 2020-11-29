package date

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Date  time.Time
		Valid bool
	}
	type args struct {
		bs []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    *Date
	}{
		{
			fields: fields{
				Date:  time.Time{},
				Valid: false,
			},
			args: args{
				bs: []byte{},
			},
			want: &Date{
				Date:  time.Time{},
				Valid: false,
			},
			wantErr: true,
		},
		{
			fields: fields{
				Date:  time.Time{},
				Valid: false,
			},
			args: args{
				bs: []byte(""),
			},
			want: &Date{
				Date:  time.Time{},
				Valid: false,
			},
			wantErr: true,
		},
		{
			fields: fields{
				Date:  time.Time{},
				Valid: false,
			},
			args: args{
				bs: []byte("1234"),
			},
			want: &Date{
				Date:  time.Time{},
				Valid: false,
			},
			wantErr: true,
		},
		{
			fields: fields{
				Date:  time.Time{},
				Valid: false,
			},
			args: args{
				bs: []byte("\"\""),
			},
			want: &Date{
				Date:  time.Time{},
				Valid: false,
			},
			wantErr: false,
		},
		{
			fields: fields{
				Date:  time.Time{},
				Valid: false,
			},
			args: args{
				bs: []byte("\"12345-01-01\""),
			},
			want: &Date{
				Date:  time.Time{},
				Valid: false,
			},
			wantErr: true,
		},
		{
			fields: fields{
				Date:  time.Time{},
				Valid: false,
			},
			args: args{
				bs: []byte("\"2020-01-01\""),
			},
			want: &Date{
				Date:  time.Date(2020, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Date{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			if err := d.UnmarshalJSON(tt.args.bs); (err != nil) != tt.wantErr {
				t.Errorf("Date.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(d, tt.want); diff != "" {
				t.Errorf("Date.UnmarshalJSON() diff = %v", diff)
			}
		})
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	type fields struct {
		Date  time.Time
		Valid bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			fields: fields{
				Date:  time.Time{},
				Valid: false,
			},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			fields: fields{
				Date:  time.Date(12345, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			fields: fields{
				Date:  time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			want:    []byte("\"2020-01-01\""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Date{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			got, err := d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Date.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.UnmarshalJSON() diff = %v", diff)
			}
		})
	}
}
