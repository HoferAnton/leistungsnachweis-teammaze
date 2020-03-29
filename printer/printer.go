package printer

import (
	"errors"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

const wall = "\u2588\u2588"
const noWall = "  "
const post = "\u2588\u2588"
const perimeter = "\u2591\u2591"
const cellNormal = "  "
const cellUp = "\u2191\u2191"
const cellDown = "\u2193\u2193"
const cellTower = "\u2193\u2191"
const nl = "\n"

const gridStep = 1

func Print2D(lab common.Labyrinth) (string, error) {
	if lab == nil {
		return "", errors.New("got nil")
	}

	_, _, maxZ := lab.GetMaxLocation().As3DCoordinates()

	var out string

	for z := uint(0); z <= maxZ; z++ {
		floor, _ := interpretFloor(lab, z)
		out = floor + out

		if z+1 <= maxZ {
			out = nl + out
		}
	}

	return out, nil
}

func interpretFloor(lab common.Labyrinth, z uint) (string, error) {
	if lab == nil {
		return "", errors.New("got nil")
	}

	maxX, maxY, maxZ := lab.GetMaxLocation().As3DCoordinates()

	if z > maxZ {
		return "", errors.New("z out of range")
	}

	const (
		one = 1
		two = 2
	)

	horizontalPerimeterLength := (maxX+one)*two + one

	out := horizontalPerimeter(horizontalPerimeterLength)

	for y := uint(0); y <= maxY; y++ {
		line, _ := interpretLine(lab, y, z)
		out = line + out
	}

	return horizontalPerimeter(horizontalPerimeterLength) + out, nil
}

func horizontalPerimeter(length uint) string {
	var out string

	for x := uint(0); x < length; x++ {
		out += perimeter
	}

	return out + nl
}

func interpretLine(lab common.Labyrinth, y uint, z uint) (string, error) {
	if lab == nil {
		return "", errors.New("got nil")
	}

	maxX, maxY, maxZ := lab.GetMaxLocation().As3DCoordinates()

	if z > maxZ {
		return "", errors.New("z out of range")
	}

	if y > maxY {
		return "", errors.New("y out of range")
	}

	out := perimeter

	for x := uint(0); x <= maxX; x++ {
		hasCeiling := !lab.IsConnected(common.NewLocation(x, y, z), common.NewLocation(x, y, z+gridStep))
		hasFloor := !lab.IsConnected(common.NewLocation(x, y, z), common.NewLocation(x, y, z-gridStep))

		switch {
		case hasCeiling && hasFloor:
			out += cellNormal
		case hasCeiling:
			out += cellDown
		case hasFloor:
			out += cellUp
		default:
			out += cellTower
		}

		if x+1 <= maxX && lab.IsConnected(common.NewLocation(x, y, z), common.NewLocation(x+gridStep, y, z)) {
			out += noWall
		} else if x+1 <= maxX {
			out += wall
		}
	}

	out += perimeter + nl

	if y > 0 {
		out += perimeter

		for x := uint(0); x <= maxX; x++ {
			if lab.IsConnected(common.NewLocation(x, y, z), common.NewLocation(x, y-gridStep, z)) {
				out += noWall
			} else {
				out += wall
			}

			if x+1 <= maxX {
				out += post
			}
		}

		out += perimeter + nl
	}

	return out, nil
}
