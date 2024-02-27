package utils

import (
	"strconv"
	"strings"
)

// HexToRGBA converts hex color to RGBA
func HexToRGB(hex string) (r, g, b uint32, err error) {
	hex = strings.TrimPrefix(hex, "#")

	val, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return
	}

	r = uint32((val >> 16))
	g = uint32((val >> 8) & 0xFF)
	b = uint32(val & 0xFF)

	return
}
