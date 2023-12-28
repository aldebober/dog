// Package dog provides converts of human to dog life in years
// It tooks human years and converts to dog's
package dog

// CONV const is converter number contsant
const CONV = 7

// Years returns int years of dog's life 
func Years(y int) int {
	return y * CONV
}
