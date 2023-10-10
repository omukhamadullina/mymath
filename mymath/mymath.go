package mymath

import "math"

var Version string = "v1.0.0"

func Abs(x float64) float64 {
	return math.Abs(x)
}
func Acos(x float64) float64 {
	return math.Acos(x)
}
func Acosh(x float64) float64 {
	return math.Acosh(x)
}
func Asin(x float64) float64 {
	return math.Asin(x)
}
func Atan(x float64) float64 {
	return math.Atan(x)
}
func Tan(x float64) float64 {
	return math.Tan(x)
}
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}
func Ceil(x float64) float64 {
	return math.Ceil(x)
}
func Floor(x float64) float64 {
	return math.Floor(x)
}
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}
func Max(x, y float64) float64 {
	return math.Max(x, y)
}
func Min(x, y float64) float64 {
	return math.Min(x, y)
}
func Round(x float64) float64 {
	return math.Round(x)
}
func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
func Mod(x, y float64) float64 {
	return math.Mod(x, y)
}
