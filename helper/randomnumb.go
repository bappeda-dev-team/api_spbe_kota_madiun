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

func GenerateRandomKodeDomain() string {
	rand.NewSource(time.Now().UnixNano())

	min := 10000
	max := 99999

	randomInt := rand.Intn(max-min) + min

	randomKodeDomain := "d-" + strconv.Itoa(randomInt)

	return randomKodeDomain
}
