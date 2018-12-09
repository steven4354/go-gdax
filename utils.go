package gdax

import (
	"strconv"
)

func StringToFloat (numStr string) float64 {
	floatType, _ := strconv.ParseFloat(numStr, 64)
	return floatType
}