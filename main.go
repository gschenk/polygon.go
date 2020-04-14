package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"polyGo/extrema"
	g "polyGo/geobjects"
	"polyGo/vector"
)

func main() {

	// Data input receives a collection of points in the
	// Cartesian plane with `x` and `y` coordinates.
	// Each line piped into STDIN is a complete collection.

	// Points are provided as a Json of the form:
	// [{"x": 1e-3, "y":-2.235}, {"x": 42, "y":-13}]
	var record []vector.PointXY
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
		points := vector.PointsToVecs(record)

		// Akl-Toussaint heuristics:
		// Quickly find a polygon that inscribes the convex hull polygon
		// by finding points with extreme coordinates along cardinal and diagonal
		// Cartesian coordinates. Henceforth the polygon constructed such is named
		// Akl-Toussaint polygon (ATP)

		// find extreme values from slice of position vectors
		extPoints := vector.RemoveDuplicates(extrema.FindValues(points[0], points))

		// create ATP
		atPoly := g.NewPoly(
			g.MapNodeFunToVecs(g.NewNode, extPoints),
		)
		fmt.Println(atPoly)

		// find geometric centre of ATP

		// Modified Gift-Wrap algorithm where only points outside the AKpoly are considered

	}

}
