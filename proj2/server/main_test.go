package main

import (
	"math/rand"
	"sort"
	"testing"
)

//func TestDeleteDuplicates(t *testing.T) {
//	for _, test := range []struct {
//		xs     []int
//		want   int
//		update []int
//	}{
//		{nil, 0, nil},
//		{[]int{1}, 1, []int{1}},
//		{[]int{1, 1}, 1, []int{1}},
//		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
//		{[]int{1, 1, 2, 2, 3, 3}, 3, []int{1, 2, 3}},
//		{[]int{2, 3, 5, 5, 7, 11, 11, 11, 13}, 6, []int{2, 3, 5, 7, 11, 13}},
//	} {
//		xs := append([]int(nil), test.xs...)
//		xxs, got := DeleteDuplicates(xs)
//		if got != test.want {
//			t.Errorf("DeleteDuplicates(%v) = %d; want %d", test.xs, got, test.want)
//		}
//		if !reflect.DeepEqual(xxs, test.update) {
//			t.Errorf("DeleteDuplicates(%v): updated slice %v; want %v", test.xs, xxs, test.update)
//		}
//	}
//}

func benchDeleteDuplicates(b *testing.B, size int) {
	b.StopTimer()
	ints := rand.New(rand.NewSource(int64(size))).Perm(size)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		DeleteDuplicates2(ints)
		b.StopTimer()
	}
}

func benchDeleteDuplicates2(b *testing.B, size int) {
	b.StopTimer()
	ints := rand.New(rand.NewSource(int64(size))).Perm(size)
	sort.Ints(ints)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		DeleteDuplicates(ints)
		b.StopTimer()
	}
}

func BenchmarkDeleteDuplicates1e3(b *testing.B) { benchDeleteDuplicates(b, 1e3) }
func BenchmarkDeleteDuplicates1e5(b *testing.B) { benchDeleteDuplicates(b, 1e5) }
func BenchmarkDeleteDuplicates1e7(b *testing.B) { benchDeleteDuplicates(b, 1e7) }

func BenchmarkDeleteDuplicates21e3(b *testing.B) { benchDeleteDuplicates2(b, 1e3) }
func BenchmarkDeleteDuplicates21e5(b *testing.B) { benchDeleteDuplicates2(b, 1e5) }
func BenchmarkDeleteDuplicates21e7(b *testing.B) { benchDeleteDuplicates2(b, 1e7) }