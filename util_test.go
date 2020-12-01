package date

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDate_AddMonths(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		date Date
		args args
		want Date
	}{
		{
			date: MustParse("2020-01-01"), args: args{0},
			want: MustParse("2020-01-01"),
		},
		{
			date: MustParse("2020-01-01"), args: args{1},
			want: MustParse("2020-02-01"),
		},
		{
			date: MustParse("2020-01-15"), args: args{1},
			want: MustParse("2020-02-15"),
		},
		{
			date: MustParse("2020-01-31"), args: args{1},
			want: MustParse("2020-02-29"),
		},
		{
			date: MustParse("2020-01-01"), args: args{2},
			want: MustParse("2020-03-01"),
		},
		{
			date: MustParse("2020-01-15"), args: args{2},
			want: MustParse("2020-03-15"),
		},
		{
			date: MustParse("2020-01-31"), args: args{2},
			want: MustParse("2020-03-31"),
		},
		{
			date: MustParse("2020-01-01"), args: args{12},
			want: MustParse("2021-01-01"),
		},
		{
			date: MustParse("2020-01-15"), args: args{12},
			want: MustParse("2021-01-15"),
		},
		{
			date: MustParse("2020-02-29"), args: args{12},
			want: MustParse("2021-02-28"),
		},
		{
			date: MustParse("2020-02-29"), args: args{37},
			want: MustParse("2023-03-29"),
		},

		// num < 0
		{
			date: MustParse("2020-01-01"), args: args{-1},
			want: MustParse("2019-12-01"),
		},
		{
			date: MustParse("2020-01-15"), args: args{-1},
			want: MustParse("2019-12-15"),
		},
		{
			date: MustParse("2020-01-31"), args: args{-1},
			want: MustParse("2019-12-31"),
		},
		{
			date: MustParse("2020-01-01"), args: args{-2},
			want: MustParse("2019-11-01"),
		},
		{
			date: MustParse("2020-01-15"), args: args{-2},
			want: MustParse("2019-11-15"),
		},
		{
			date: MustParse("2020-01-31"), args: args{-2},
			want: MustParse("2019-11-30"),
		},
		{
			date: MustParse("2020-01-01"), args: args{-12},
			want: MustParse("2019-01-01"),
		},
		{
			date: MustParse("2020-01-15"), args: args{-12},
			want: MustParse("2019-01-15"),
		},
		{
			date: MustParse("2020-02-29"), args: args{-12},
			want: MustParse("2019-02-28"),
		},
		{
			date: MustParse("2020-02-29"), args: args{-37},
			want: MustParse("2017-01-29"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := tt.date
			got := date.AddMonths(tt.args.num)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.AddMonths() = diff %v", diff)
			}
		})
	}
}

func TestDate_DiffMonths(t *testing.T) {
	type args struct {
		ref Date
	}
	tests := []struct {
		name string
		date Date
		args args
		want int
	}{
		{
			date: MustParse("2020-01-01"),
			args: args{MustParse("2020-01-01")},
			want: 1,
		},
		{
			date: MustParse("2020-01-01"),
			args: args{MustParse("2020-01-31")},
			want: 1,
		},
		{
			date: MustParse("2020-01-01"),
			args: args{MustParse("2020-02-01")},
			want: 2,
		},
		{
			date: MustParse("2020-01-01"),
			args: args{MustParse("2020-02-02")},
			want: 2,
		},
		{
			date: MustParse("2020-02-01"),
			args: args{MustParse("2020-01-02")},
			want: -1,
		},
		{
			date: MustParse("2020-02-02"),
			args: args{MustParse("2020-01-01")},
			want: -2,
		},
		{
			date: MustParse("2020-01-01"),
			args: args{MustParse("2020-12-31")},
			want: 12,
		},
		{
			date: MustParse("2020-01-01"),
			args: args{MustParse("2021-01-31")},
			want: 13,
		},
		{
			date: MustParse("2020-01-01"),
			args: args{MustParse("2021-02-01")},
			want: 14,
		},
		{
			date: MustParse("2020-01-01"),
			args: args{MustParse("2019-01-02")},
			want: -12,
		},
		{
			date: MustParse("2020-02-01"),
			args: args{MustParse("2019-01-02")},
			want: -13,
		},
		{
			date: MustParse("2020-02-01"),
			args: args{MustParse("2019-01-01")},
			want: -14,
		},
		{
			date: MustParse("2020-01-31"),
			args: args{MustParse("2020-02-28")},
			want: 1,
		},
		{
			date: MustParse("2020-01-31"),
			args: args{MustParse("2020-02-29")},
			want: 2,
		},
		{
			date: MustParse("2020-01-31"),
			args: args{MustParse("2020-03-30")},
			want: 2,
		},
		{
			date: MustParse("2020-02-28"),
			args: args{MustParse("2020-03-31")},
			want: 2,
		},
		{
			date: MustParse("2020-03-31"),
			args: args{MustParse("2020-02-29")},
			want: -2,
		},
		{
			date: MustParse("2020-03-31"),
			args: args{MustParse("2020-03-01")},
			want: -1,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := tt.date
			got := date.DiffMonths(tt.args.ref)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Date.DiffMonths() #%d has diff %v", i, diff)
			}
		})
	}
}
