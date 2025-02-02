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

const maxrecursion = 1e5
const fewpoints = 2000
const debugflag = false
const debuglog = "./debug.log"

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

		// filter duplicates for small collections of points
		if len(points) < fewpoints {
			points = points.Undup()
		}

		// take care of special cases
		if len(points) < 3 {
			fmt.Println(0) //area 0
			continue
		}

		// Akl-Toussaint heuristics:
		// Quickly find a polygon that inscribes the convex hull polygon
		// by finding points with extreme coordinates along cardinal and diagonal
		// Cartesian coordinates. Henceforth the polygon constructed such is named
		// Akl-Toussaint polygon (ATP)

		// find extreme values from slice of position vectors, dump duplicates
		extPoints := extrema.FindValues(points[0], points).Undup()

		// extrema.FindValues does not find ATP for very slim shapes
		// use (expensive )second strategy:
		if len(extPoints) == 2 {
			extPoints = extrema.SecondExtremes(points, extPoints[0], extPoints[1])
		}

		// take care of degenerate cases
		if len(extPoints) < 3 {
			fmt.Println(0) //area 0
			continue
		}

		// create ATP
		atPoly := geo.NewPoly(
			geo.NewNodes(extPoints),
		)
		//fmt.Println("Akl-Toussaint polygon area", atPoly.Area)

		// Modified Gift-Wrap algorithm where only points outside the AKpoly are considered

		// find points outside of each edge of the ATP and store them with the edge
		atPoly.FindOutsidePoints(points)

		// run a recursive chain search nor CHP nodes outside each ATP edge
		chpNodes := chainSearch(atPoly.Edges)

		// construct CHP
		chPoly := geo.NewPoly(
			chpNodes,
		)
		//fmt.Println("Complex hull polygon area", chPoly.Area)

		// output CHP area
		if !chPoly.IsComplete {
			debugWrite("Complex hull poly not complete.", chPoly, text)
		}
		fmt.Println(chPoly.Area)

	}
}
