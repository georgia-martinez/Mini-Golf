package main

import (
	"image/color"
	"strconv"
	"math"
)

func GetDistance(x0 float64, y0 float64, x1 float64, y1 float64) float64 {
	return math.Sqrt(math.Pow(x1 - x0, 2) + math.Pow(y1 - y0, 2))
}

func GetClosestNumber(a float64, b float64, c float64) float64 {
	if math.Abs(c - a) < math.Abs(c - b) {
		return a
	} else {
        return b
	}
}

func GetRGBColor(R int64, G int64, B int64) color.RGBA {
	
	r := DecToHex(R)
	g := DecToHex(G)
	b := DecToHex(B)

	return color.RGBA{r, g, b, 0xff}
}

func DecToHex(num int64) uint8 {

	hexString := strconv.FormatInt(num, 16)
	hexNum, _ := strconv.ParseInt(hexString, 16, 64)

	return uint8(hexNum)
}