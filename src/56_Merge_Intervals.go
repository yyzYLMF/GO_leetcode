package main

import "fmt"
import "sort"

type Interval struct {
	Start int
	End   int
}

type IntervalArr []Interval

func (arr IntervalArr) Len() int {
	return len(arr)
}

func (arr IntervalArr) Less(i, j int) bool {
	return arr[i].Start < arr[j].Start
}

func (arr IntervalArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(IntervalArr(intervals))
	ret := []Interval{}
	ret = append(ret, intervals[0])
	pos := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start > ret[pos].End {
			ret = append(ret, intervals[i])
			pos++
		} else {
			if ret[pos].End < intervals[i].End {
				ret[pos].End = intervals[i].End
			}
		}
	}
	return ret
}

func main() {
	input := []Interval{
		Interval{
			Start: 1,
			End:   3,
		},
		Interval{
			Start: 2,
			End:   6,
		},
	}

	fmt.Println(merge(input))
	return
}
