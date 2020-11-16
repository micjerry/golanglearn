package main

import (
	"fmt"
	"time"
	"gopl.io/ch12/format"
)

type Movie struct {
	Title, Subtitle    string
	Year               int
	Color              bool
	Actor              map[string]string
	Oscars             []string
	Sequel             *string
}

var sequel = "Hello"

var strangelove = Movie {
	Title:       "Dr Strangelove",
	Subtitle:    "How I learned to Stop worrying and Love the Bomb",
	Year:        1964,
	Color:       false,
	Actor:       map[string]string{
			"Dr. Strangelove":           "Peter Sellers",
			"Grp. Capt Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":      "Peter Sellers",
			"Gen. Buck Turgidson":       "George C. Scott",
			"Brig. Gen. Jack D. Ripper": "Sterling Hayden",
			`Maj. T.J. "King" Kong`:     "Slim Pickens",
	},
	Oscars: []string{
			"Best Actor",
			"Best Adapted Screenplay",
			"Best Director",
			"Best Picture",
	},
	Sequel: &sequel,
}

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	
	fmt.Println(format.Any(x))
	fmt.Println(format.Any(d))
	fmt.Println(format.Any([]int64{x}))
	fmt.Println(format.Any([]time.Duration{d}))
	
	format.Display("strangelove", strangelove)
}