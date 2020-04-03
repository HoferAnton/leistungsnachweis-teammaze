package printer

import (
	"errors"

	"github.com/ob-algdatii-20ss/leistungsnachweis-teammaze/common"
)

const (
	wall       = "██"
	noWall     = "  "
	noWallPath = "░░"
	post       = "██"
	perimeter  = "▒▒"

	cellNormal = "  "
	cellDown   = "↓↓"
	cellUp     = "↑↑"
	cellUpDown = "↑↓"

	pathNormal        = "░░"
	pathDownCould     = "░↓"
	pathDownGo        = "▼▼"
	pathUpCould       = "↑░"
	pathUpGo          = "▲▲"
	pathUpDownCould   = "⇅░"
	pathUpCouldDownGo = "↑▼"
	pathUpGoDownCould = "▲↓"
	pathUpDownGo      = "▲▼"

	nl = "\n"

	gridStep = 1
)

/**
@param path can be nil
*/
func Print2D(lab common.Labyrinth, path []common.Location) (string, error) {
	if lab == nil {
		return "", errors.New("no maze given")
	}

	_, _, maxZ := lab.GetMaxLocation().As3DCoordinates()

	var out string

	for z := uint(0); z <= maxZ; z++ {
		out = interpretFloor(lab, z, path) + out

		if z+1 <= maxZ {
			out = nl + out
		}
	}

	return out, nil
}

func interpretFloor(lab common.Labyrinth, z uint, path []common.Location) string {
	if lab == nil {
		panic("require a maze")
	}

	maxX, maxY, maxZ := lab.GetMaxLocation().As3DCoordinates()

	if z > maxZ {
		panic("z out of range")
	}

	const (
		one = 1
		two = 2
	)

	horizontalPerimeterLength := (maxX+one)*two + one

	out := horizontalPerimeter(horizontalPerimeterLength)

	for y := uint(0); y <= maxY; y++ {
		out = interpretLine(lab, y, z, path) + out
	}

	return horizontalPerimeter(horizontalPerimeterLength) + out
}

func horizontalPerimeter(length uint) string {
	var out string

	for x := uint(0); x < length; x++ {
		out += perimeter
	}

	return out + nl
}

func interpretLine(lab common.Labyrinth, y uint, z uint, path []common.Location) string {
	if lab == nil {
		panic("got nil")
	}

	maxX, maxY, maxZ := lab.GetMaxLocation().As3DCoordinates()

	if z > maxZ || y > maxY {
		panic("z or y out of range")
	}

	out := perimeter

	for x := uint(0); x <= maxX; x++ {
		common.NewLocation(x, y, z)
		out += interpretCell(lab, common.NewLocation(x, y, z), path)

		if x+1 <= maxX {
			if lab.IsConnected(common.NewLocation(x, y, z), common.NewLocation(x+gridStep, y, z)) {
				if contains(path, common.NewLocation(x, y, z)) && contains(path, common.NewLocation(x+gridStep, y, z)) {
					out += noWallPath
				} else {
					out += noWall
				}
			} else {
				out += wall
			}
		}
	}

	out += perimeter + nl

	//line 0 of a floor has no wall underneath it so we exit early
	if y == 0 {
		return out
	}

	out += perimeter

	for x := uint(0); x <= maxX; x++ {
		if lab.IsConnected(common.NewLocation(x, y, z), common.NewLocation(x, y-gridStep, z)) {
			if contains(path, common.NewLocation(x, y, z)) && contains(path, common.NewLocation(x, y-gridStep, z)) {
				out += noWallPath
			} else {
				out += noWall
			}
		} else {
			out += wall
		}

		if x+1 <= maxX {
			out += post
		}
	}

	out += perimeter + nl

	return out
}

func interpretCell(lab common.Labyrinth, position common.Location, path []common.Location) string {
	if lab == nil || position == nil {
		panic("got nil")
	}

	x, y, z := position.As3DCoordinates()

	isOnPath := contains(path, common.NewLocation(x, y, z))

	hasCeiling := !lab.IsConnected(common.NewLocation(x, y, z), common.NewLocation(x, y, z+gridStep))
	hasFloor := !lab.IsConnected(common.NewLocation(x, y, z), common.NewLocation(x, y, z-gridStep))

	if isOnPath {
		goUp := !hasCeiling && contains(path, common.NewLocation(x, y, z+gridStep))
		goDown := !hasFloor && contains(path, common.NewLocation(x, y, z-gridStep))

		return selectPath(hasFloor, hasCeiling, goDown, goUp)
	}

	switch {
	case hasCeiling && hasFloor:
		return cellNormal
	case hasCeiling:
		return cellDown
	case hasFloor:
		return cellUp
	default:
		return cellUpDown
	}
}

func selectPath(hasFloor bool, hasCeiling bool, goDown bool, goUp bool) string {
	switch {
	case goUp && goDown:
		return pathUpDownGo
	case goUp && !goDown && !hasFloor:
		return pathUpGoDownCould
	case !goUp && goDown && !hasCeiling:
		return pathUpCouldDownGo
	case !goUp && !goDown && !hasCeiling && !hasFloor:
		return pathUpDownCould
	case hasFloor && goUp:
		return pathUpGo
	case !hasCeiling && hasFloor && !goUp:
		return pathUpCould
	case hasCeiling && goDown:
		return pathDownGo
	case hasCeiling && !hasFloor && !goDown:
		return pathDownCould
	default:
		return pathNormal
	}
}

func contains(l []common.Location, e common.Location) bool {
	for _, s := range l {
		if s.Compare(e) {
			return true
		}
	}

	return false
}
