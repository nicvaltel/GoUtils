package main

import (
	"fmt"

	_ "github.com/nicvaltel/GoUtils/filter"
	_ "github.com/nicvaltel/GoUtils/log"
	_ "github.com/nicvaltel/GoUtils/maybe"
	"github.com/nicvaltel/GoUtils/test"
	_ "github.com/nicvaltel/GoUtils/types"
)

func main() {
	fmt.Println("GoUtils Testing...")
	test.RunAllTests()
}
