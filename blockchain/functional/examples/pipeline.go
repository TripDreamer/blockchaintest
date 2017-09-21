package main

import (
	"fmt"
	f "github.com/izqui/functional"
	"strings"
)

type Band struct {
	Name    string
	Country string
	Active  bool
}

type BandFunc func(b Band) Band

func capitalizeName(b Band) Band {

	a := b
	a.Name = strings.ToUpper(a.Name)

	return a
}

func setCountryToSpain(b Band) Band {

	a := b
	a.Country = "Spain"

	return a
}

func pipeline(fs []BandFunc, bs []Band) []Band {

	return f.DoMap(func(b Band) Band {

		return f.Reduce(func(b Band, f BandFunc) Band {

			return f(b)

		}, fs, b).(Band)

	}, bs).([]Band)
}

func main() {

	bands := []Band{Band{Name: "beatles", Country: "UK"}, Band{Name: "coldplay", Country: "UK"}}
	newBands := pipeline([]BandFunc{capitalizeName, setCountryToSpain}, bands)
	fmt.Println(newBands)

	newBands2 := f.Pipeline([]func(b Band) Band{capitalizeName, setCountryToSpain}, bands)
	fmt.Println(newBands2)
}
