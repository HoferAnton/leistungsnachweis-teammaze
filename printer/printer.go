package printer

import (
	"fmt"
	. "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

const wall = "\u2588\u2588"
const perimeter = "\u2591\u2591"
const cellNormal = "  "
const cellUp = "\u2191\u2191"
const cellDown = "\u2193\u2193"
const cellTower = "\u2193\u2191"

func Print2D(lab Labyrinth) {
	maxX, maxY, maxZ := lab.GetMaxLocation().As3DCoordinates()
	for z := maxZ; z <= maxZ; z-- { //layers top to bottom     // and yes i need to check for underflow since i need the 0

		fmt.Println("") // new line between floors

		for x := uint(0); x <= ((maxX+1)*2+1)-1; x++ {
			fmt.Print(perimeter)
		}
		fmt.Println("")

		for y := maxY; y <= maxY; y-- { //lines top to bottom     // and yes i need to check for underflow since i need the 0
			printline(lab, y, z)
		}

		for x := uint(0); x <= ((maxX+1)*2+1)-1; x++ {
			fmt.Print(perimeter)
		}
		fmt.Println("")
	}
}

func printline(lab Labyrinth, y uint, z uint) {

	maxX, _, _ := lab.GetMaxLocation().As3DCoordinates()

	fmt.Print(perimeter)

	for x := uint(0); x <= maxX; x++ {

		ceiling := !lab.IsConnected(NewLocation(x, y, z), NewLocation(x, y, z+1))
		floor := !lab.IsConnected(NewLocation(x, y, z), NewLocation(x, y, z-1))
		if ceiling && floor {
			fmt.Print(cellNormal)
		} else if ceiling {
			fmt.Print(cellDown)
		} else if floor {
			fmt.Print(cellUp)
		} else {
			fmt.Print(cellTower)
		}

		if x+1 <= maxX {
			if lab.IsConnected(NewLocation(x, y, z), NewLocation(x+1, y, z)) {
				fmt.Print("  ")
			} else {
				fmt.Print(wall)
			}
		}

	}

	fmt.Println(perimeter)

	if y > 0 {
		fmt.Print(perimeter)
		for x := uint(0); x <= maxX; x++ {
			if lab.IsConnected(NewLocation(x, y, z), NewLocation(x, y-1, z)) {
				fmt.Print("  ")
			} else {
				fmt.Print(wall)
			}
			if x+1 <= maxX {
				fmt.Print(wall) // pfosten
			}
		}

		fmt.Println(perimeter)
	}

}
