package main

import (
	"fmt"

	"github.com/jgarc172/gocsprog/lab6-2/ratio"
)

func main() {
	sdtv := ratio.New(4, 3)
	hd := ratio.New(16, 9)
	aspect := new(ratio.Ratio)
	aspect = aspect.SumInst(sdtv, hd)
	fmt.Println(aspect)
	aspect2 := ratio.StaticSum(sdtv, hd)
	fmt.Println(aspect2)
	aspect3 := sdtv.Sum(hd)
	fmt.Println(aspect3)
}
