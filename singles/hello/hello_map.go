package main

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lmicroseconds)
}

func createAndInitMap() {
	log.Println("-------------------------------------------------------------------------")
	log.Println("| creation and initialization of maps")
	log.Println("-------------------------------------------------------------------------")

	func() {
		var dict map[string]string // nil
		log.Printf("just declared map: %v == nil", dict)
	}()

	func() {
		dict := make(map[string]int)
		log.Printf("created by make(): %v", dict)
	}()

	func() {
		dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
		log.Printf("created and initialized by a map literal: %v", dict)
	}()

	// Won't be compiled
	// The map key can be a value from any built-in or struct type as log as the value can be used in an expression with the == operator.
	// Slices, functions, and struct types that contain slices can't be used as map keys.
	//func() {
	//	dict := map[[]string]int{}
	//	log.Printf("%v", dict)
	//}()

	func() {
		dict := map[string][]string{"Korea": []string{"Seoul"}, "USA": []string{"Washington", "New York"}}
		log.Printf("slices as map values: %v", dict)
	}()

	log.Println("")
}

func manipulateMap() {
	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| modify elements of a map by [] operator")
		log.Println("-------------------------------------------------------------------------")

		colors := map[string]string{}
		log.Printf("map: %v, size=%d", colors, len(colors))

		colors["Red"] = "#da1337"
		log.Printf("after modifying by []: %v, size=%d", colors, len(colors))

		log.Println("")
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| Retrieving a value from a map and testing existence")
		log.Println("-------------------------------------------------------------------------")

		colors := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
		log.Printf("map: %v, size=%d", colors, len(colors))

		orange, orangeExists := colors["Orange"]
		log.Printf("Orange in the map? %t, it's %v", orangeExists, orange)

		blue, blueExists := colors["Blue"]
		log.Printf("Blue in the map? %t, is it zero string? %t", blueExists, blue == "")

		log.Println()
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| Deleting a value from a map")
		log.Println("-------------------------------------------------------------------------")

		colors := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
		log.Printf("map: %v, size=%d", colors, len(colors))

		delete(colors, "Orange")

		orange, orangeExists := colors["Orange"]
		log.Printf("Orange in the map? %t, is it zero string? %t", orangeExists, orange == "")

		log.Println()
	}()
}

func iterateMap() {
	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| iteration by for-range")
		log.Println("-------------------------------------------------------------------------")

		colors := map[string]string{
			"AliceBlue":   "#f0f8ff",
			"Coral":       "#ff7F50",
			"DarkGray":    "#a9a9a9",
			"ForestGreen": "#228b22",
		}
		log.Printf("map: %v, size=%d", colors, len(colors))

		for key, value := range colors {
			log.Printf("\t%s -> %s", key, value)
		}

		log.Println("")
	}()
}

func passMap() {
	removeFunc := func(target map[string]string, keys ...string) {
		for _, k := range keys {
			delete(target, k)
		}
	}

	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	log.Printf("map: %v, size=%d", colors, len(colors))

	// it's safe for missing key(Blue) to be deleted
	removeFunc(colors, "AliceBlue", "DarkGray", "Blue")

	log.Printf("after removing some elements by a function call: %v, size=%d", colors, len(colors))
}

func main() {
	createAndInitMap()
	manipulateMap()
	iterateMap()
	passMap()
}
