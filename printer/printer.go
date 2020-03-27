package printer

import . "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"

const wall = "\u2588\u2588"
const noWall = "  "
const post = "\u2588\u2588"
const perimeter = "\u2591\u2591"
const cellNormal = "  "
const cellUp = "\u2191\u2191"
const cellDown = "\u2193\u2193"
const cellTower = "\u2193\u2191"
const nl = "\n"

func Print2D(lab Labyrinth) string {
	if lab == nil {
		return ""
	}
	_, _, maxZ := lab.GetMaxLocation().As3DCoordinates()

	var out string
	for z := uint(0); z <= maxZ; z++ {
		out = interpretFloor(lab, z) + out
		if z+1 <= maxZ {
			out = nl + out
		}
	}
	return out
}

func interpretFloor(lab Labyrinth, z uint) string {
	maxX, maxY, _ := lab.GetMaxLocation().As3DCoordinates()

	out := horizontalPerimeter((maxX+1)*2 + 1)

	for y := uint(0); y <= maxY; y++ {
		out = interpretLine(lab, y, z) + out
	}

	return horizontalPerimeter((maxX+1)*2+1) + out
}

func horizontalPerimeter(length uint) string {
	var out string
	for x := uint(0); x < length; x++ {
		out += perimeter
	}
	return out + nl
}

func interpretLine(lab Labyrinth, y uint, z uint) string {

	maxX, _, _ := lab.GetMaxLocation().As3DCoordinates()

	out := perimeter

	for x := uint(0); x <= maxX; x++ {

		hasCeiling := !lab.IsConnected(NewLocation(x, y, z), NewLocation(x, y, z+1))
		hasFloor := !lab.IsConnected(NewLocation(x, y, z), NewLocation(x, y, z-1))
		if hasCeiling && hasFloor {
			out += cellNormal
		} else if hasCeiling {
			out += cellDown
		} else if hasFloor {
			out += cellUp
		} else {
			out += cellTower
		}

		if x+1 <= maxX {
			if lab.IsConnected(NewLocation(x, y, z), NewLocation(x+1, y, z)) {
				out += noWall
			} else {
				out += wall
			}
		}

	}

	out += perimeter + nl

	if y > 0 {
		out += perimeter
		for x := uint(0); x <= maxX; x++ {
			if lab.IsConnected(NewLocation(x, y, z), NewLocation(x, y-1, z)) {
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

	return out
}
