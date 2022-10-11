package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func OrderSn() string {
	r := rand.Intn(10000)
	h := fmt.Sprintf("%04d", r)
	i := time.Now().Unix() + time.Now().UnixNano()
	s := strconv.FormatInt(i, 10)
	format := time.Now().Format("20060102150405")
	return format + s[len(s)-4:] + h
}
