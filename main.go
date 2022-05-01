package main

import (
	"fmt"

	"github.com/kodernubie/keireport"
)

func main() {

	rpt, err := keireport.LoadFromFile("simple.krpt")

	if err == nil {

		rpt.GenToFile("simple.pdf")
	} else {

		fmt.Println("Error loading template file :", err.Error())
	}
}
