package main

import (
	"github.com/paulsmith/gogeos/geos"
	"strconv"
)

func main() {
	serveHttp()
}

func testCoordinatePresence(x, y, polygyon string) bool {
	fx, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return false
	}

	fy, err := strconv.ParseFloat(y, 64)
	if err != nil {
		return false
	}

	point, err := geos.NewPoint(geos.NewCoord(fx, fy))
	if err != nil {
		return false
	}

	boundary, err := geos.FromWKT(polygyon)
	if err != nil {
		return false
	}

	result, err := boundary.Contains(point)
	if err != nil {
		return false
	}

	return result
}
