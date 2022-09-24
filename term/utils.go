package term

import (
	"math"
	"strconv"
)

func getBoostValue(boostSymbol string) float64 {
	if len(boostSymbol) == 0 || boostSymbol == "^" {
		return 1.0
	} else {
		var res, _ = strconv.ParseFloat(boostSymbol[1:], 64)
		return res
	}
}

func getFuzziness(fuzzySymbol string) int {
	if fuzzySymbol == "~" {
		// default fuzziness
		return -1
	} else {
		var v, _ = strconv.ParseFloat(fuzzySymbol[1:], 64)
		return int(math.Round(v))
	}
}
