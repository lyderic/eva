package main

import (
	"fmt"
)

func add(name, value string) (err error) {
	construction()
	return
}

func get(name string) (value string) {
	construction()
	return
}

func list() {
	construction()
}

func del(name string) (err error) {
	construction()
	return
}

func construction() {
	fmt.Println("In construction...")
}
