package utils

import (
	"math/rand"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

func GenerateUniqueKey() uint {
	mu.Lock()
	defer mu.Unlock()

	var length = 7
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	key := uint(0)
	for i := 0; i < length; i++ {
		key = key*10 + uint(seededRand.Intn(9)) + 1
	}

	return key
}
