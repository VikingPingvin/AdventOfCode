package main

import (
	"reflect"
	"strings"
	"testing"
)

const (
	testinput = `467..114..
		...*......
		..35..633.
		......#...
		617*......
		.....+.58.
		..592.....
		......755.
		...$.*....
		.664.598..`
)

func textToTrimmedSlice(t *testing.T, input string) []string {
	t.Helper()
	// trim test data
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.ReplaceAll(input, "\r", "")
	res := strings.Split(input, "\n")
	return res
}

func Test_solveP1(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{
				data: testinput,
			},
			want: 4361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// trim test data
			// data := strings.ReplaceAll(tt.args.data, "\t", "")
			// data = strings.ReplaceAll(data, "\r", "")
			// dataSplit := strings.Split(data, "\n")
			result, _ := solveP1(textToTrimmedSlice(t, tt.args.data))

			if result != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, result)
			}
		})
	}

}

func Test_getPartNumbers(t *testing.T) {
	type args struct {
		pos  int
		line []byte
	}
	tests := []struct {
		name              string
		args              args
		wantBytePositions []int
		wantNumbers       int
	}{
		{
			name: "test-1",
			args: args{
				pos:  4,
				line: []byte("..406."),
			},
			wantBytePositions: []int{2, 3, 4},
			wantNumbers:       406,
		},
		{
			name: "test-2",
			args: args{
				pos:  2,
				line: []byte("..406."),
			},
			wantBytePositions: []int{2, 3, 4},
			wantNumbers:       406,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytePositions, gotNumbers := getPartNumbers(tt.args.pos, tt.args.line)
			if !reflect.DeepEqual(gotBytePositions, tt.wantBytePositions) {
				t.Errorf("getPartNumbers() gotBytePositions = %v, want %v", gotBytePositions, tt.wantBytePositions)
			}
			if gotNumbers != tt.wantNumbers {
				t.Errorf("getPartNumbers() gotNumbers = %v, want %v", gotNumbers, tt.wantNumbers)
			}
		})
	}
}

func Test_solveP2(t *testing.T) {
	type args struct {
		symbolCollection map[int]map[int][]int
		lines            []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test-1",
			args: args{
				symbolCollection: map[int]map[int][]int{
					1: {
						3: {
							467, 35,
						},
					},
					8: {
						5: {
							755, 598,
						},
					},
				},
				lines: textToTrimmedSlice(t, testinput),
			},
			want: 467835,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveP2(tt.args.symbolCollection, tt.args.lines); got != tt.want {
				t.Errorf("solveP2() = %v, want %v", got, tt.want)
			}
		})
	}
}
