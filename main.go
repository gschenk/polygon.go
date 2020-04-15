package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"polyGo/extrema"
	"polyGo/geo"
	"polyGo/vector"
)

const maxrecursion = 1000

func main() {

	// Data input receives a collection of points in the
	// Cartesian plane with `x` and `y` coordinates.
	// Each line piped into STDIN is a complete collection.

	// Points are provided as a Json of the form:
	// [{"x": 1e-3, "y":-2.235}, {"x": 42, "y":-13}]
	var record vector.Points
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal([]byte(text), &record)
		if err != nil {
			log.Fatal(err)
		}

		// assign data from json to a slice of vectors
		points := record.Vecs()

		// Akl-Toussaint heuristics:
		// Quickly find a polygon that inscribes the convex hull polygon
		// by finding points with extreme coordinates along cardinal and diagonal
		// Cartesian coordinates. Henceforth the polygon constructed such is named
		// Akl-Toussaint polygon (ATP)

		// find extreme values from slice of position vectors
		extPoints := vector.RemoveDuplicates(
			extrema.FindValues(points[0], points),
		)

		// create ATP
		atPoly := geo.NewPoly(
			geo.NewNodes(extPoints),
		)

		// find geometric centre of ATP
		fmt.Println("Akl-Toussaint polygon area", atPoly.Area)

		// Modified Gift-Wrap algorithm where only points outside the AKpoly are considered

		// find points outside of each edge of the ATP and store them with the edge
		atPoly.FindOutsidePoints(points)

		// run a recursive chain search nor CHP nodes outside each ATP edge
		chpNodes := chainSearch(atPoly.Edges)
		fmt.Println("ATP", atPoly.Nodes.Ids())
		fmt.Println("CHP", chpNodes.Ids())

	}
}
