package panics

import (
	"fmt"
)

func ExamplePanicAndHandle() {
	defer func() {

		if v := recover(); v != nil {
			fmt.Println("Востановление после паники")
		}
	}()

	var notInitializedMap map[string]string
	notInitializedMap["key"] = "value"
}
