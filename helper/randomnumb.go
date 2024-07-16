package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomKode() string {
	rand.NewSource(time.Now().UnixNano())

	min := 1000
	max := 9999

	randomInt := rand.Intn(max-min) + min

	randomKodeProsesBisnis := strconv.Itoa(randomInt)

	return randomKodeProsesBisnis
}
