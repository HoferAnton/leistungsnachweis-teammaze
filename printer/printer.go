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

func Print2D(lab Labyrinth) string {
	var out string
	maxX, maxY, maxZ := lab.GetMaxLocation().As3DCoordinates()
	for z := maxZ; z <= maxZ; z-- { //layers top to bottom     // and yes i need to check for underflow since i need the 0

		out += "\n" // free line between floors

		for x := uint(0); x <= ((maxX+1)*2+1)-1; x++ {
			out += perimeter
		}
		out += "\n"

		for y := maxY; y <= maxY; y-- { //lines top to bottom     // and yes i need to check for underflow since i need the 0
			out += InterpretLine(lab, y, z)
		}

		for x := uint(0); x <= ((maxX+1)*2+1)-1; x++ {
			out += perimeter
		}
		out += "\n"
	}
	return out
}

func InterpretLine(lab Labyrinth, y uint, z uint) string {

	maxX, _, _ := lab.GetMaxLocation().As3DCoordinates()

	out := perimeter

	for x := uint(0); x <= maxX; x++ {

		ceiling := !lab.IsConnected(NewLocation(x, y, z), NewLocation(x, y, z+1))
		floor := !lab.IsConnected(NewLocation(x, y, z), NewLocation(x, y, z-1))
		if ceiling && floor {
			out += cellNormal
		} else if ceiling {
			out += cellDown
		} else if floor {
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

	out += perimeter + "\n"

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

		out += perimeter + "\n"
	}

	return out
}

//TODO: dynamic newline ?
