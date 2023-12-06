package main

import (
	"reflect"
	"testing"
)

func Test_splitLine(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name               string
		args               args
		wantCardNum        int
		wantWinningNumbers []int
		wantHandNumbers    []int
	}{
		{
			name: "test1",
			args: args{
				input: "Card   1: 26 36 90  2 75 32  3 21 59 18 | 47 97 83 82 43  7 61 73 57  2 67 31 69 11 44 38 23 52 10 21 45 36 86 49 14",
			},
			wantCardNum:        1,
			wantWinningNumbers: []int{26, 36, 90, 2, 75, 32, 3, 21, 59, 18},
			wantHandNumbers:    []int{47, 97, 83, 82, 43, 7, 61, 73, 57, 2, 67, 31, 69, 11, 44, 38, 23, 52, 10, 21, 45, 36, 86, 49, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCardNum, gotWinningNumbers, gotHandNumbers := splitLine(tt.args.input)
			if gotCardNum != tt.wantCardNum {
				t.Errorf("splitLine() gotCardNum = %v, want %v", gotCardNum, tt.wantCardNum)
			}
			if !reflect.DeepEqual(gotWinningNumbers, tt.wantWinningNumbers) {
				t.Errorf("splitLine() gotWinningNumbers = %v, want %v", gotWinningNumbers, tt.wantWinningNumbers)
			}
			if !reflect.DeepEqual(gotHandNumbers, tt.wantHandNumbers) {
				t.Errorf("splitLine() gotHandNumbers = %v, want %v", gotHandNumbers, tt.wantHandNumbers)
			}
		})
	}
}

func Test_stringSliceToIntSlice(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				input: []string{"1", "5", "15"},
			},
			want:    []int{1, 5, 15},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := stringSliceToIntSlice(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringSliceToIntSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stringSliceToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveP1(t *testing.T) {
	type args struct {
		cards *CardCollection
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				cards: &CardCollection{
					cards: map[int]Card{
						1: {
							winningNumbers: []int{9000, 150},
							handNumbers:    []int{1, 2, 3, 4, 5, 6, 7, 9000},
						},
					},
				},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				cards: &CardCollection{
					cards: map[int]Card{
						1: {
							winningNumbers: []int{9000, 150},
							handNumbers:    []int{1, 2, 3, 9000, 9000, 150},
						},
					},
				},
			},
			want: 4,
		},
		{
			name: "test-multiple-cards",
			args: args{
				cards: &CardCollection{
					cards: map[int]Card{
						1: {
							winningNumbers: []int{9000, 150},
							handNumbers:    []int{1, 2, 3, 9000, 9000, 150},
						},
						2: {
							winningNumbers: []int{9000, 150},
							handNumbers:    []int{1, 2, 3, 9000, 9000, 150},
						},
					},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveP1(tt.args.cards); got != tt.want {
				t.Errorf("solveP1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveP2(t *testing.T) {
	type args struct {
		cards *CardCollection
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		cards: &CardCollection{
		// 			cards: map[int]Card{
		// 				1: {
		// 					winningNumbers: []int{1, 2},
		// 					handNumbers:    []int{1, 2, 9, 9, 9},
		// 				},
		// 				2: {
		// 					winningNumbers: []int{9, 9},
		// 					handNumbers:    []int{1, 1, 1, 1},
		// 				},
		// 				3: {
		// 					winningNumbers: []int{9, 9},
		// 					handNumbers:    []int{1, 1, 1, 1},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: 5,
		// },
		{
			name: "test-originaltestinput",
			// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
			// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
			// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
			// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
			// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
			// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
			args: args{
				cards: &CardCollection{
					cards: map[int]Card{
						1: {
							winningNumbers: []int{41, 48, 83, 86, 17},
							handNumbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
						},
						2: {
							winningNumbers: []int{13, 32, 20, 16, 61},
							handNumbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
						},
						3: {
							winningNumbers: []int{1, 21, 53, 59, 44},
							handNumbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
						},
						4: {
							winningNumbers: []int{41, 92, 73, 84, 69},
							handNumbers:    []int{59, 84, 76, 51, 58, 5, 54, 83},
						},
						5: {
							winningNumbers: []int{87, 83, 26, 28, 32},
							handNumbers:    []int{88, 30, 70, 12, 93, 22, 82, 36},
						},
						6: {
							winningNumbers: []int{31, 18, 13, 56, 72},
							handNumbers:    []int{74, 77, 10, 23, 35, 67, 36, 11},
						},
					},
				},
			},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveP2(tt.args.cards); got != tt.want {
				t.Errorf("solveP2() = %v, want %v", got, tt.want)
			}
		})
	}
}
