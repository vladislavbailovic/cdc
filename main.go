package main

import (
	"fmt"
	"komunalno/cdc/pkg/codec"
)

func main() {
	x := codec.Html("T&#x25;e&#x3C;s&#x3E;tI&#x22;n&#x27;gString")
	fmt.Println(x.Decode())
}
