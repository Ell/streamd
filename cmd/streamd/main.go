package main

import (
	"fmt"

	"github.com/ell/streamd"
)

func main() {
	point, _ := streamd.GetCursorPos()
	fmt.Printf("point: %+v\n", point)
}
