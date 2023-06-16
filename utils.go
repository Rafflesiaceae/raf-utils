package main

import (
	"fmt"
	"os"
	"strconv"
)

func fromInt(arg string) int {
	i, err := strconv.Atoi(arg)
	checkError(err)
	return i
}

func die(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func checkErrorDie(err error) {
	if err != nil {
		die(err.Error())
	}
}
