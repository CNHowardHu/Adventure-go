package Random

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandInt(LBound, Len int) int {
	return LBound + r.Intn(Len)
}
