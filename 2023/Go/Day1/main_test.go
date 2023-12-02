package main

import (
	"testing"
)

func Test_convertLineToDigits(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal string",
			args: args{
				line: "one54",
			},
			want: "154",
		},
		{
			name: "abnormal string",
			args: args{
				line: "on54",
			},
			want: "54",
		},
		{
			name: "long 1",
			args: args{
				line: "35smnsnzxmdjtsns6sevenonethree",
			},
			want: "356713",
		},
		{
			name: "long 2",
			args: args{
				line: "ncxxlsqdkvc8fiverslzqtzhzltcmbkthreelkjjckxsvljvs",
			},
			want: "853",
		},
		{
			name: "short",
			args: args{
				line: "v4",
			},
			want: "4",
		},
		{
			name: "what",
			args: args{
				line: "eightthree",
			},
			want: "83",
		},
		{
			name: "thefuck",
			args: args{
				line: "eighthree",
			},
			want: "83",
		},
		{
			name: "isthis",
			args: args{
				line: "eightwo",
			},
			want: "82",
		},
		{
			name: "comeon",
			args: args{
				line: "coneightfivedfkqrfjcckghzsrtrc9sevenone1",
			},
			want: "1859711",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertLineToDigits(tt.args.line); got != tt.want {
				t.Errorf("convertLineToDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCalibrationValues(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				line: "v4",
			},
			want: 8,
		},
		{
			name: "test2",
			args: args{
				line: "7seven8threeeight",
			},
			want: 15,
		},
		{
			name: "test3",
			args: args{
				line: "ptnqxxf1two",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCalibrationValues(tt.args.line); got != tt.want {
				t.Errorf("getCalibrationValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
