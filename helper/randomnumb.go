package helper

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomKodeProsesBisnis() string {
	rand.NewSource(time.Now().UnixNano())

	min := 1
	max := 99999999999

	randomInt := rand.Intn(max-min) + min

	randomKodeProsesBisnis := "pb-" + strconv.Itoa(randomInt)

	return randomKodeProsesBisnis
}

func GenerateRandomKodeLayananSPBE() string {
	rand.NewSource(time.Now().UnixNano())

	min := 1
	max := 99999999999

	randomInt := rand.Intn(max-min) + min

	randomKodeProsesBisnis := "l-" + strconv.Itoa(randomInt)

	return randomKodeProsesBisnis
}

func GenerateRandomKodeDatanInformasi() string {
	rand.NewSource(time.Now().UnixNano())

	min := 1
	max := 99999999999

	randomInt := rand.Intn(max-min) + min

	randomKodeProsesBisnis := "di-" + strconv.Itoa(randomInt)

	return randomKodeProsesBisnis
}

func GenerateRandomKodeAplikasi() string {
	rand.NewSource(time.Now().UnixNano())

	min := 1
	max := 99999999999

	randomInt := rand.Intn(max-min) + min

	randomKodeProsesBisnis := "a-" + strconv.Itoa(randomInt)

	return randomKodeProsesBisnis
}

func GenerateRandomKodeDomain() string {
	rand.NewSource(time.Now().UnixNano())

	min := 1
	max := 99999999999

	randomInt := rand.Intn(max-min) + min

	randomKodeDomain := "d-" + strconv.Itoa(randomInt)

	return randomKodeDomain
}
