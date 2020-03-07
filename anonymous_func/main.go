package main

import (
	"fmt"
	"go-mockery-mock-sample/anonymous_func/funcs"
	"log"
)

func main() {
	log.Print("start")
	ex := funcs.NewExample(funcs.NewExampleCalled())

	if result, err := ex.ExampleAnonymousFunc(); err != nil {
		log.Print("error!")
	} else {
		log.Print(fmt.Sprintf("result: %s", result))
	}
	log.Print("end")
}
