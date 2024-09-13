package helper

import (
	"log"
	"strconv"
)

func ConvertStringsToInterfaces(strings []string) []interface{} {
	interfaces := make([]interface{}, len(strings))
	for i, s := range strings {
		interfaces[i] = s
	}
	return interfaces
}

// StringToInt mengkonversi string menjadi int
// Jika terjadi error, fungsi ini akan mengembalikan 0
func ConvertStringToInt(s string) int {
	if s == "" {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		return 0
	}
	return i
}
