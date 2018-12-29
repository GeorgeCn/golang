package main

import "image/color"

type ColorPoint struct {
	*Point
	Color color.RGBA
}

