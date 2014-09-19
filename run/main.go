package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ostlerc/ballclock/baller"
	//"github.com/davecheney/profile"
)

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()

	if len(os.Args) < 2 {
		fmt.Println("Too few arguments")
		return
	}

	var clocks [10]uint8
	clockv := int8(0)
	found := false

	for i := 1; i < len(os.Args); i++ {
		v, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Printf("Invalid argument '%s'\n", os.Args[i])
			return
		}
		//stopping condition
		if v == 0 {
			found = true
			break
		}

		if v < 27 || v > 127 {
			fmt.Println("arguments out of range. Should be within [27,127].")
			return
		}
		if clockv == 10 {
			fmt.Println("Too many arguments, only 10 allowed.")
			return
		}

		clocks[clockv] = uint8(v)
		clockv++
	}

	if !found {
		fmt.Println("Missing end argument 0")
		return
	}

	for i := int8(0); i < clockv; i++ {
		days := baller.EvalBallClockV6(clocks[i])
		fmt.Printf("%d ball cycle after %d days.\n", clocks[i], days)
	}
}
