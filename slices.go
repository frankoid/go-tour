package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	println(dx, dy)
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
	}

	for y := range pic {
		row := pic[y]
		for x := range row {
			//row[x] = uint8((x+y)/2)
			row[x] = uint8(x ^ y)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
