package term

import (
	"strconv"
)

type BoostValue float64

func (b BoostValue) Float() float64 {
	return float64(b)
}

type Fuzziness float64

func (f Fuzziness) Float() float64 {
	return float64(f)
}

var AutoFuzzy Fuzziness = -1      // only ~
var DefaultBoost BoostValue = 1.0 // no boost symbol

var NoFuzzy Fuzziness = 0.0
var NoBoost BoostValue = 0.0

func getBoostValue(boostSymbol string) BoostValue {
	if len(boostSymbol) == 0 || boostSymbol == "^" {
		// default boost
		return DefaultBoost
	} else {
		var v, _ = strconv.ParseFloat(boostSymbol[1:], 64)
		return BoostValue(v)
	}
}

func getFuzzyValue(fuzzySymbol string) Fuzziness {
	if fuzzySymbol == "~" {
		// auto fuzziness
		return AutoFuzzy
	} else {
		var v, _ = strconv.ParseFloat(fuzzySymbol[1:], 64)
		return Fuzziness(v)
	}
}
