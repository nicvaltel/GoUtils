package test

import (
	"fmt"
	"reflect"

	"github.com/nicvaltel/GoUtils/filter"
	"github.com/nicvaltel/GoUtils/maybe"
)

func testStep[A any](testName string, value A, etalon A) {
	result := reflect.DeepEqual(value, etalon)
	fmt.Printf("checking %20s: %v\n", testName, result)
	if !result {
		fmt.Printf("\t result = %v\n", value)
		fmt.Printf("\t etalon = %v\n", etalon)
	}
}

func testFilter(ints []int, strs []string) {
	checkFilter1 := filter.Filter(func(n int) bool { return (n > 1 && n <= 4) }, ints)
	checkFilterGauge1 := []int{2, 3, 3}
	testStep("filter 1", checkFilter1, checkFilterGauge1)

	checkFilter2 := filter.Filter(func(s string) bool { return (len(s) > 2 && len(s) <= 5) }, strs)
	checkFilterGauge2 := []string{"Hello", "dear"}
	testStep("filter 2", checkFilter2, checkFilterGauge2)

	checkFilter3 := filter.TakeFirst(func(s string) bool { return (len(s) < 3) }, strs)
	checkFilterGauge3 := maybe.Pure("my")
	testStep("filter 3", checkFilter3, checkFilterGauge3)

	checkFilter4 := filter.TakeN(4, func(x int) bool { return (x > 5) }, ints)
	checkFilterGauge4 := []int{7, 8, 12, 55}
	testStep("filter 4", checkFilter4, checkFilterGauge4)

	checkFilter5 := filter.TakeN(4, func(x int) bool { return (x > 10) }, ints)
	checkFilterGauge5 := []int{12, 55, 50}
	testStep("filter 5", checkFilter5, checkFilterGauge5)

}

func RunAllTests() {
	ints := []int{1, 2, 3, 7, 8, 12, -3, 0, 55, 3, 1, 1, 50}
	strs := []string{"Hello", "my", "dear", "friend"}

	testFilter(ints, strs)

}
