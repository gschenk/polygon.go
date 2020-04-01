package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type Point struct{
    X float64
    Y float64
}

        // [{"x": 1e-3, "y":-2.235}, {"x": 42, "y":-13}]
func main() {
    var record []Point
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
        fmt.Printf("x: %f, y: %f\n", record[0].X, record[0].Y)

    }
}
