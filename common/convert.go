package common

import "strconv"

func StringSliceToIntSlice(str []string) []int {
	interval := make([]int, 0)
	for _, v := range str {
		i, _ := strconv.Atoi(v)
		interval = append(interval, i)
	}
	return interval
}
func StringMinuteToIntSecond(str []string) []int {
	interval := make([]int, 0)
	for _, v := range str {
		i, _ := strconv.Atoi(v)
		interval = append(interval, i*60)
	}
	return interval
}
