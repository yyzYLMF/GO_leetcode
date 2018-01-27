package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		intervals = append(intervals, newInterval)
		return intervals
	}
	pos := 0
	ret := []Interval{}
	for ; pos < len(intervals); pos++ {
		if newInterval.Start < intervals[pos].Start {
			break
		} else {
			ret = append(ret, intervals[pos])
		}
	}
	if pos == 0 || newInterval.Start > intervals[pos-1].End {
		ret = append(ret, newInterval)
	} else {
		if ret[pos-1].End < newInterval.End {
			ret[pos-1].End = newInterval.End
		}
	}
	retLen := len(ret)
	for i := pos; i < len(intervals); i++ {
		if ret[retLen-1].End < intervals[i].Start {
			ret = append(ret, intervals[i:]...)
			return ret
		} else {
			if intervals[i].End > ret[retLen-1].End {
				ret[retLen-1].End = intervals[i].End
			}
		}
	}
	return ret
}

func main() {
	input := []Interval{}
	item := Interval{
		Start: 1,
		End:   2,
	}
	input = append(input, item)
	item = Interval{
		Start: 3,
		End:   4,
	}
	fmt.Println(insert(input, item))
	return
}
