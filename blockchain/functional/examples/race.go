package main

import (
	"fmt"
	f "github.com/izqui/functional"
	"github.com/izqui/helpers"
	"time"
)

type Race struct {
	Time int
	Cars []Car
}

type Car struct {
	Position int
}

func draw(r Race) {

	drawing := f.Map(func(car Car) string {

		return f.Reduce(func(a string, i int) string {

			return a + "-"

		}, make([]int, car.Position), "").(string)

	}, r.Cars).([]string)

	fmt.Println("............................")
	f.Map(func(a string) interface{} {

		fmt.Println(a)
		return a
	}, drawing)
}

func carState(cars []Car) []Car {

	return f.Map(func(c Car) Car {

		if helpers.RandomInt(0, 3) == 0 {

			return Car{Position: c.Position + 1}
		}

		return Car{Position: c.Position}

	}, cars).([]Car)
}

func raceLoop(r Race) Race {

	return Race{Time: r.Time - 1, Cars: carState(r.Cars)}
}

func race(r Race) {

	if r.Time > 0 {

		draw(r)
		time.Sleep(100 * time.Millisecond)
		race(raceLoop(r))

	} else {

		fmt.Println("Race finished")
	}
}

func main() {

	r := Race{Time: 100, Cars: []Car{Car{}, Car{}, Car{}}}
	race(r)
}
