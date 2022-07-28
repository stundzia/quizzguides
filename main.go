package main

import (
	"fmt"
	"github.com/stundzia/quizzguides/guides"
)

func main() {
	fmt.Println("Quizzguides")
	guides.MuxVsRWMux()
}
