package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapCopy(t *testing.T) {
	map1 := map[int]string{
		1: "asd",
		2: "wakawaka",
	}

	map2 := make(map[int]string, len(map1))

	for k, v := range map1 {
		map2[k] = v
	}
	map2[1] = "wakawaka"

	fmt.Printf("Map1: %v\n", map1)
	fmt.Printf("Map2: %v\n", map2)

}

func Test_solveP1_Example(t *testing.T) {
	lines, err := readFileToLines("testinput.txt")
	if err != nil {
		t.Errorf("failed reading file: %v", err)
	}
	want := 35

	got := solveP1(lines)

	if got[0] != want {

		t.Errorf("gotSlice: %v, got %v, want%v\n", got, got[0], want)
	}
}

func Test_solveP2_Example(t *testing.T) {
	lines, err := readFileToLines("testinput.txt")
	if err != nil {
		t.Errorf("failed reading file: %v", err)
	}
	want := 46

	got := solveP2(lines)

	if got != want {

		t.Errorf("got %v, want%v\n", got, want)
	}
}

func Test_getSeeds(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				line: "seeds: 1234 56 78 910",
			},
			want: []int{1234, 56, 78, 910},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumbers(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSeeds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeElementFromSlice(t *testing.T) {
	type args struct {
		input      []interface{}
		startIndex int
		endIndex   int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "test1",
			args: args{
				input:      []interface{}{1, 2, 3, 4},
				startIndex: 2,
				endIndex:   2,
			},
			want: []interface{}{1, 2, 4},
		},
		{
			name: "test2",
			args: args{
				input:      []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 124},
				startIndex: 0,
				endIndex:   7,
			},
			want: []interface{}{9, 124},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElementFromSlice(tt.args.input, tt.args.startIndex, tt.args.endIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElementFromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_modifyRangeSlice(t *testing.T) {
	type args struct {
		slice     []int
		index     int
		newValues []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				slice:     []int{1, 2, 3, 4},
				index:     0,
				newValues: []int{10, 19, 20, 20},
			},
			want: []int{10, 19, 20, 20, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modifyRangeSlice(tt.args.slice, tt.args.index, tt.args.newValues); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("modifyRangeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSeedRangeLocationMap(t *testing.T) {
	type args struct {
		seeds         *RangeSeeds
		mapCollection map[int]SourceDestinationMap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test - sourcerange within seedrange",
		// 	args: args{
		// 		seeds: &RangeSeeds{
		// 			Seeds: []SeedRange{
		// 				SeedRange{
		// 					min: 46,
		// 					max: 56,
		// 				},
		// 				SeedRange{
		// 					min: 78,
		// 					max: 80,
		// 				},
		// 			},
		// 		},
		// 		mapCollection: map[int]SourceDestinationMap{
		// 			0: SourceDestinationMap{
		// 				name: "first",
		// 				maps: []Mapping{
		// 					Mapping{
		// 						src:  50,
		// 						dest: 52,
		// 						rng:  2,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: 46,
		// },
		// {
		// 	name: "test - seedrange within sourcerange 1",
		// 	args: args{
		// 		seeds: &RangeSeeds{
		// 			Seeds: []SeedRange{
		// 				SeedRange{
		// 					min: 79,
		// 					max: 92,
		// 				},
		// 			},
		// 		},
		// 		mapCollection: map[int]SourceDestinationMap{
		// 			0: SourceDestinationMap{
		// 				name: "first",
		// 				maps: []Mapping{
		// 					Mapping{
		// 						src:  50,
		// 						dest: 98,
		// 						rng:  2,
		// 					},
		// 					Mapping{
		// 						src:  50,
		// 						dest: 52,
		// 						rng:  48,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: 81,
		// },
		// {
		// 	name: "test - seedrange within sourcerange 2",
		// 	args: args{
		// 		seeds: &RangeSeeds{
		// 			Seeds: []SeedRange{
		// 				SeedRange{
		// 					min: 79,
		// 					max: 92,
		// 				},
		// 				SeedRange{
		// 					min: 51,
		// 					max: 55,
		// 				},
		// 			},
		// 		},
		// 		mapCollection: map[int]SourceDestinationMap{
		// 			0: SourceDestinationMap{
		// 				name: "first",
		// 				maps: []Mapping{
		// 					Mapping{
		// 						src:  50,  // 50-51 (+10)
		// 						dest: 60,
		// 						rng:  2,
		// 					},
		// 					Mapping{
		// 						src:  45, // 45-48 (-4)
		// 						dest: 49,
		// 						rng:  4,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: 52,
		// },
		// {
		// 	name: "test - seedrange bottom collision 1",
		// 	args: args{
		// 		seeds: &RangeSeeds{
		// 			Seeds: []SeedRange{
		// 				SeedRange{
		// 					min: 10,
		// 					max: 20,
		// 				},
		// 				SeedRange{
		// 					min: 250,
		// 					max: 350,
		// 				},
		// 			},
		// 		},
		// 		mapCollection: map[int]SourceDestinationMap{
		// 			0: SourceDestinationMap{
		// 				name: "first",
		// 				maps: []Mapping{
		// 					Mapping{
		// 						src:  8, // 8-17 (+10) --> [10 17 18 20 250 350]
		// 						dest: 18, // ^-> [20 27 18 20 250 350]
		// 						rng:  10,
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// 	want: 18,
		// },
		{
			name: "test - seedrange contains source range",
			args: args{
				seeds: &RangeSeeds{
					Seeds: []SeedRange{
						SeedRange{
							min: 10,
							max: 20,
						},
						// SeedRange{
						// 	min: 21,
						// 	max: 30,
						// },
					},
				},
				mapCollection: map[int]SourceDestinationMap{
					0: SourceDestinationMap{
						name: "first",
						maps: []Mapping{
							Mapping{
								src:  10, // 10-19 (+10) --> [10 19 20 20 21 30]
								dest: 20, // ^-> [20 29 20 20 21 30]
								rng:  10,
							},
						},
					},
				},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSeedRangeLocationMap(tt.args.seeds, tt.args.mapCollection); got != tt.want {
				t.Errorf("getSeedRangeLocationMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
