package nulldate

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/lovung/date"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Date  date.Date
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
		want    *NullDate
	}{
		{
			fields: fields{
				Date:  date.NewZero(),
				Valid: false,
			},
			args: args{
				bs: []byte{},
			},
			want: &NullDate{
				Date:  date.NewZero(),
				Valid: false,
			},
			wantErr: true,
		},
		{
			fields: fields{
				Date:  date.NewZero(),
				Valid: false,
			},
			args: args{
				bs: []byte(""),
			},
			want: &NullDate{
				Date:  date.NewZero(),
				Valid: false,
			},
			wantErr: true,
		},
		{
			fields: fields{
				Date:  date.NewZero(),
				Valid: false,
			},
			args: args{
				bs: []byte("1234"),
			},
			want: &NullDate{
				Date:  date.NewZero(),
				Valid: false,
			},
			wantErr: true,
		},
		{
			fields: fields{
				Date:  date.NewZero(),
				Valid: false,
			},
			args: args{
				bs: []byte("\"\""),
			},
			want: &NullDate{
				Date:  date.NewZero(),
				Valid: false,
			},
			wantErr: false,
		},
		{
			fields: fields{
				Date:  date.NewZero(),
				Valid: false,
			},
			args: args{
				bs: []byte("\"12345-01-01\""),
			},
			want: &NullDate{
				Date:  date.NewZero(),
				Valid: false,
			},
			wantErr: true,
		},
		{
			fields: fields{
				Date:  date.NewZero(),
				Valid: false,
			},
			args: args{
				bs: []byte("\"2020-01-01\""),
			},
			want: &NullDate{
				Date:  date.New(time.Date(2020, time.Month(1), 1, 0, 0, 0, 0, time.UTC)),
				Valid: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &NullDate{
				Date:  tt.fields.Date,
				Valid: tt.fields.Valid,
			}
			if err := d.UnmarshalJSON(tt.args.bs); (err != nil) != tt.wantErr {
				t.Errorf("NullDate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(d, tt.want); diff != "" {
				t.Errorf("NullDate.UnmarshalJSON() diff = %v", diff)
			}
		})
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	type fields struct {
		Date  date.Date
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
				Date:  date.NewZero(),
				Valid: false,
			},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			fields: fields{
				Date:  date.New(time.Date(12345, time.January, 1, 0, 0, 0, 0, time.UTC)),
				Valid: true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			fields: fields{
				Date:  date.New(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)),
				Valid: true,
			},
			want:    []byte("\"2020-01-01\""),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := NullDate{
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
