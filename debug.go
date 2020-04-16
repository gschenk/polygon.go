package main

import (
	"fmt"
	"os"
)

func debugWrite(text string, something interface{}, raw string) {
	isErr := func(e error) {
		if e != nil {
			panic(e)
		}
	}
	if debugflag {
		fmt.Fprintf(os.Stderr, "Warning: %s", text)

		file, err := os.Create(debuglog)
		isErr(err)

		defer file.Close()

		fmt.Fprintf(file, "Error: %s\nStruct %+v\n%s", text, something, raw)

	}
}
