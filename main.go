package main

import . "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
import . "github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/printer"

func main() {
	lab := NewLabyrinth(NewLocation(5, 3, 3))
	lab.Connect(NewLocation(1, 1, 2), NewLocation(1, 1, 1))
	lab.Connect(NewLocation(1, 1, 2), NewLocation(1, 1, 3))
	lab.Connect(NewLocation(1, 1, 2), NewLocation(1, 0, 2))
	lab.Connect(NewLocation(1, 1, 2), NewLocation(1, 2, 2))
	lab.Connect(NewLocation(1, 1, 2), NewLocation(0, 1, 2))
	lab.Connect(NewLocation(1, 1, 2), NewLocation(2, 1, 2))
	Print2D(lab)
}
