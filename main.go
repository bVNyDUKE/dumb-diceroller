package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}

	arg := strings.Split(os.Args[1], "d")

	if len(arg) != 2 {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}

	number, err := strconv.Atoi(arg[0])
	if err != nil {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}

	dice, err := strconv.Atoi(arg[1])
	if err != nil {
		fmt.Println("Invalid argument")
		os.Exit(1)
	}

	res := 0
	for i := 1; i <= number; i++ {
		roll := rand.Intn(dice) + 1
		fmt.Printf("%v \n", roll)
		res += roll
	}

	fmt.Printf("You rolled %v \n", res)
}
