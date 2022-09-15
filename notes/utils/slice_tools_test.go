package utils

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test_intersect_slice",
			args: args{
				[]string{"1", "2", "3", "6", "8"},
				[]string{"2", "3", "5", "0"},
			},
			want: []string{"2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasIntersect(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test_has_intersect_slice",
			args: args{
				[]string{"1", "2", "3", "6", "8"},
				[]string{"2", "5", "4", "0"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasIntersect(tt.args.slice1, tt.args.slice2); got != tt.want {
				t.Errorf("HasIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubset(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test_subset_slice",
			args: args{
				[]string{"2", "3"},
				[]string{"1", "2", "3", "6", "8"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subset(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsValue(t *testing.T) {
	type args struct {
		slice  []string
		target string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "test_contains_value_slice",
			args: args{
				[]string{"2"},
				"2",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsValue(tt.args.slice, tt.args.target); got != tt.want {
				t.Errorf("ContainsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test_diff_slice_1",
			args: args{
				slice1: []string{"2", "5", "7", "3", "6", "8", "1"},
				slice2: []string{"3", "6", "8", "1"},
			},
			want: []string{"5", "2", "7"},
		},
		{
			name: "test_diff_slice_2",
			args: args{
				slice1: []string{"2", "5", "7"},
				slice2: []string{"3", "6", "2", "5", "8", "1", "4"},
			},
			want: []string{"3", "6", "8", "1", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceV2(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test_diff_slice_1",
			args: args{
				a: []string{"2", "5", "7", "3", "6", "8", "1"},
				b: []string{"3", "6", "8", "1"},
			},
			want: []string{"2", "5", "7"},
		},
		{
			name: "test_diff_slice_2",
			args: args{
				a: []string{"2", "5", "7"},
				b: []string{"3", "6", "2", "5", "8", "1", "4"},
			},
			want: []string{"7", "3", "6", "8", "1", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceV2(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DifferenceV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifferenceV3(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test_diff_slice_1",
			args: args{
				a: []string{"2", "5", "7", "3", "6", "8", "1"},
				b: []string{"3", "6", "8", "1"},
			},
			want: []string{"2", "5", "7"},
		},
		{
			name: "test_diff_slice_2",
			args: args{
				b: []string{"3", "6", "2", "5", "8", "1", "4"},
				a: []string{"2", "5", "7"},
			},
			want: []string{"7", "3", "6", "8", "1", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceV3(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DifferenceV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test_union_1",
			args: args{
				slice2: []string{"2", "5", "7"},
				slice1: []string{"3", "6"},
			},
			want: []string{"3", "6", "2", "5", "7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRepeat(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test_union_1",
			args: args{
				arr: []string{"2", "2", "3", "3", "5", "7"},
			},
			want: []string{"2", "3", "5", "7"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveRepeat(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRepeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
