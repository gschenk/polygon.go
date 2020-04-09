package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"polyGo/extrema"
	"polyGo/tools"
	"polyGo/vector"
)

// [{"x": 1e-3, "y":-2.235}, {"x": 42, "y":-13}]
func main() {
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
		points := tools.PointsToVecs(record)

		// find extreme values from :
		extPoints := tools.RemoveDuplicates(extrema.FindValues(points[0], points))
		fmt.Println(extPoints)

	}

}
