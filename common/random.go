package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func RandFloat64(min, max float64) float64 {
	n := len(strings.Split(fmt.Sprintf("%v", min), ".")[1])
	res := min + rand.Float64()*(max-min)
	float, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(n)+"f", res), 10)
	return float
}
