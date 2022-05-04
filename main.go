package main

import (
	"fmt"
	"komunalno/cdc/pkg/codec"
)

func main() {
	x := codec.Slug("wh_a__t")
	fmt.Println(x.Encode())
}
