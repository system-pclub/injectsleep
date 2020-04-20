package injectsleep

import (
	"math/rand"
	"time"
)

func SleepD () {
	num := rand.NormFloat64() * 1.0 + 0

	if num <= 0 {
		return
	}

	time.Sleep(time.Duration(num) * time.Second)
}


func Sleep(m float64) {
	num := rand.NormFloat64() * 1.0 + m

	if num <= 0 {
		return
	}

	time.Sleep(time.Duration(num) * time.Second)
}

func SleepMS (m float64, std float64) {
	num := rand.NormFloat64() * std + m

	if num <= 0 {
		return
	}

	time.Sleep(time.Duration(num) * time.Second)
}


/*
func main() {
	i := 0

	numSum := float64(0)

	for i < 100 {
		num := rand.NormFloat64() * 2.0 + 3
		numSum += num

		fmt.Print(num)
		fmt.Print(" ")

		i ++
	}

	fmt.Println()
	fmt.Println(numSum/100.0)
}*/
