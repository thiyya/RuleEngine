package utils

import (
	"math"
	"reflect"
	"strings"
)

func ConvertCommaToDot(str string) string {
	return strings.Replace(str, ",", ".", -1)
}

func ConvertDotToComma(str string) string {
	return strings.Replace(str, ".", ",", -1)
}

func Map(in interface{}, fn func(interface{}) interface{}) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		out[i] = fn(val.Index(i).Interface())
	}
	return out
}

func Reduce(in interface{}, memo interface{}, fn func(interface{}, interface{}) interface{}) interface{} {
	val := reflect.ValueOf(in)
	for i := 0; i < val.Len(); i++ {
		memo = fn(val.Index(i).Interface(), memo)
	}
	return memo
}

func Filter(in interface{}, fn func(interface{}) bool) interface{} {
	val := reflect.ValueOf(in)
	out := make([]interface{}, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		current := val.Index(i).Interface()
		if fn(current) {
			out = append(out, current)
		}
	}
	return out
}

func IsAnErrorOccured(err error) bool {
	if err == nil {
		return false
	}
	if err.Error() == "" {
		return false
	}
	return true
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
