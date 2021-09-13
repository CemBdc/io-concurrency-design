package main

import (
	"fmt"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		s := fmt.Errorf("%s: %s", msg, err)
		fmt.Printf("%s\n", s)
		panic(s)
	}
}

//badWay took 1.353974541s
//goodWay took 999.654625ms
func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}
