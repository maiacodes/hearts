package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Command line params
	rawHeartLocation := flag.String("raw", "raw.svg", "the location of the raw heart")
	heartName := flag.String("name", "lesbian", "the name of the newly minted hearts")
	hexColors := flag.String("colors", "#D62900,#FF9B55,#FFFFFF,#D461A6,#A50062", "comma seperated hex colors")
	flag.Parse()

	// Load heart into memory
	dat, err := os.ReadFile(*rawHeartLocation)
	if err != nil {
		panic(err)
	}
	fill := "DD2E44"
	heartXML := string(dat)

	// Parse colors
	colors := strings.Split(strings.ReplaceAll(*hexColors, "#", ""), ",")
	if len(colors) < 2 {
		panic("You need at least 2 stripes!")
	}

	// Create hearts
	for i, stripe := range colors {
		// Write new heart
		heartData := strings.ReplaceAll(heartXML, fill, stripe)
		name := fmt.Sprintf("%v_%v", *heartName, i+1)
		//err := os.WriteFile(name+".svg", []byte(heartData), 0644)
		err := writePNG(heartData, name)
		if err != nil {
			fmt.Println("Cannot create heart ", i, " because ", err)
		}

	}
}
