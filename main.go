package main

func main() {

	lab := generator.NewDepthFirstGenerator().GenerateLabyrinth(
		common.NewLocation(
			uint(100),
			uint(100),
			uint(10)))

	empty, _ := printer.Print2D(lab, nil)
	fmt.Printf("created  %v \n", empty != "")
	//fmt.Printf("Generated maze:\n%v", empty)

	start := common.NewLocation(0, 0, 0)
	dest := lab.GetMaxLocation()
	path := solver.RecursiveSolver(lab, start, dest, true)
	solved, _ := printer.Print2D(lab, path)
	fmt.Printf("solved  %v \n", solved != "")
	//fmt.Printf("Found path from %v to %v:\n%v \nVisualized:\n%v", start, dest, path, solved)

}
