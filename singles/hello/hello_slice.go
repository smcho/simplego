package main

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lmicroseconds)
}

func logStringSlice(prefix string, slice *[]string) {
	if *slice == nil {
		log.Printf(prefix + ": nil")
	} else if len(*slice) == 0 {
		log.Printf(prefix + ": empty, capacity=%d", cap(*slice))
	} else {
		log.Printf(prefix + ": %p->%v, length=%d, capacity=%d, isNil=%t", &(*slice)[0], *slice, len(*slice), cap(*slice), *slice == nil)
	}
}

func logIntSlice(prefix string, slice *[]int) {
	if *slice == nil {
		log.Printf(prefix + ": nil")
	} else if len(*slice) == 0 {
		log.Printf(prefix + ": empty, capacity=%d", cap(*slice))
	} else {
		log.Printf(prefix + ": %p->%v, length=%d, capacity=%d, isNil=%t", &(*slice)[0], *slice, len(*slice), cap(*slice), *slice == nil)
	}
}

func createAndInitSlice() {
	log.Println("-------------------------------------------------------------------------")
	log.Println("| creation and initialization of slices")
	log.Println("-------------------------------------------------------------------------")

	func() {
		slice := make([]string, 5)
		logStringSlice("a slice of string, created by make() with length 5", &slice)
	}()

	func() {
		slice := make([]string, 3, 5)
		logStringSlice("a slice of string, created by make() with length 3, capacity 5", &slice)
	}()

	func() {
		slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
		logStringSlice("a slice of string, initialized by a slice literal", &slice)
	}()

	func() {
		slice := []int{10, 20, 30}
		logIntSlice("a slice of integer, initialized by a slice literal", &slice)
	}()

	func() {
		slice := []string{9: ""}
		logStringSlice("a slice of string, initialized by the 10th element with an empty string", &slice)
	}()

	func() {
		array := [...]string{"apple", "google", "amazon", "facebook", "oracle"}
		log.Printf("an array(source): %p->%v, length=%d", &array, array, len(array))

		slice := array[1:4]	// a slice containing the array
		logStringSlice("slice created by slicing an array", &slice)

		whole := array[:]	// convert an array to a slice
		logStringSlice("array to slice", &whole)
	}()

	func() {
		// A nil slice -> represent a slice that doesn't exist
		// A nil slice equals(==) to nil!!!
		var slice []int
		logIntSlice("a nil slice", &slice)
	}()

	func() {
		// an empty slice -> represent an empty collection, such as when a database query returns zero results
		slice := make([]int, 0)
		logIntSlice("an empty slice", &slice)
	}()

	func() {
		// another way to create an empty slice
		slice := []int{}
		logIntSlice("another empty slice", &slice)

		slice = nil // this nil assignments is possible
		logIntSlice("after assign nil to an slice type variable", &slice)
	}()

	log.Println("")
}

func manipulateSlice() {
	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| slicing and modification by []")
		log.Println("-------------------------------------------------------------------------")

		slice := []int{10, 20, 30, 40, 50}
		logIntSlice("initial", &slice)

		slice[1] = 25
		logIntSlice("after modifying by [] operator", &slice)

		newSlice := slice[1:3]
		logIntSlice("derived slice", &newSlice)

		slice[1] = 26 // affects the slice derived from this slice <- After the slicing operation performed, the two slices share the same underlying array!
		logIntSlice("after another modifying on the original slice by [] operator", &slice)
		logIntSlice("derived slice, after modifying the original slice", &newSlice)

		newSlice[1] = 35 // affects the original slice
		logIntSlice("after modifying the derived slice by [] operator", &newSlice)
		logIntSlice("original slice, after modifying the derived slice", &slice)

		log.Println("")
	}()


	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| illegal slicing")
		log.Println("-------------------------------------------------------------------------")

		source := []int{10, 20, 30, 40, 50}
		logIntSlice("source", &source)

		possible := source[1:5] // legal, the end index(5) <= len(source)
		logIntSlice("legal slicing within the length", &possible)

		defer func() {
			if err := recover(); err != nil { // catch
				log.Printf("an expected error: %v", err)
			}
			log.Println("")
		}()

		impossible := source[1:6] // illegal, the end index(6) > len(source)
		logIntSlice("illegal slicing within the length", &impossible)
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| slicing is restricted by the capacity, not the length")
		log.Println("-------------------------------------------------------------------------")

		source := []int{10, 20, 30, 40, 50}
		logIntSlice("source", &source)

		possible := source[1:3:4] // legal, the end index(3) <= cap(source)
		logIntSlice("legal slicing within the length", &possible)

		alsoPossible := possible[0:3] // also legal, the end index(3) > cap(possible)
		logIntSlice("illegal slicing within the length", &alsoPossible)

		log.Println("")
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| append()")
		log.Println("-------------------------------------------------------------------------")

		slice := make([]int, 0, 5)
		logIntSlice("empty slice with 5 capacity", &slice)

		appended1 := append(slice, 1)
		logIntSlice("after appending an element", &appended1)
		logIntSlice("original slice", &slice)

		appended2 := append(slice, 2, 3)
		logIntSlice("after appending another element", &appended2)
		logIntSlice("first appended slice, after appending an element", &appended1)
		logIntSlice("original slice", &slice)

		log.Println("")
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| slicing and append()")
		log.Println("-------------------------------------------------------------------------")

		slice := []int{10, 20, 30, 40, 50}
		logIntSlice("initial", &slice)

		newSlice := slice[1:3]
		logIntSlice("derived slice", &newSlice)

		newSlice = append(newSlice, 60)
		logIntSlice("after appending[no growing] to the derived slice", &newSlice)
		logIntSlice("original slice, after appending[no growing] to the derived slice", &slice)

		newSlice2 := append(newSlice, 70, 80)
		logIntSlice("after appending[growing] to the derived slice (returned by append())", &newSlice2)
		logIntSlice("after appending[growing] to the derived slice", &newSlice)
		logIntSlice("original slice, after appending[growing] to the derived slice", &slice)

		newSlice2[1] = 35 // this slice was forked, so doesn't affect the original and first derived slices
		logIntSlice("after modifying the growed derived slice", &newSlice2)
		logIntSlice("the first derived slice, after modifying the growed derived slice", &newSlice)
		logIntSlice("original slice, after modifying the growed derived slice", &slice)

		log.Println("")
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| three index slices")
		log.Println("-------------------------------------------------------------------------")

		source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
		logStringSlice("source", &source)

		// This third index gives you control over the capacity of the new slice.
		// The purpose is not to increase capacity, but to restrict capacity.
		slice := source[2:3:4]
		logStringSlice("slice with third index", &slice)

		slice = append(slice, "Mango")	// affects to the original slice
		logStringSlice("after appending[no growing] to the derived slice", &slice)
		logStringSlice("original slice, after appending[growing] to the derived slice", &source)

		defer func() {
			if err := recover(); err != nil {	// catch
				log.Printf("an expected error: %v", err)
			}
			log.Println("")
		}()

		impossible := source[2:3:6]
		logStringSlice("runtime error expected", &impossible)
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| guarded slicing with third index")
		log.Println("-------------------------------------------------------------------------")

		source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
		logStringSlice("source", &source)

		guardedSlice := source[2:3:3]
		logStringSlice("slice with third index, capacity restricted", &guardedSlice)

		guardedSlice = append(guardedSlice, "Kiwi")
		logStringSlice("after appending the guarded slice", &guardedSlice)
		logStringSlice("the original slice, after appending the guarded slice", &source)

		log.Println("")
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| appending to a slice from another slice")
		log.Println("-------------------------------------------------------------------------")

		s1 := []int{1, 2}
		s2 := []int{3, 4}

		// Using ... operator, you can append all the elements of one slice into another
		s3 := append(s1, s2...)
		logIntSlice("to-slice", &s1)
		logIntSlice("from-slice", &s2)
		logIntSlice("concatenated slice", &s3)

		// if to-slice have available capacity
		s4 := make([]int, 0, 5)
		s4 = append(s4, 1, 2)
		logIntSlice("to-slice with room to append", &s4)

		s5 := append(s4, s2[0])	// s4 and s5 share the underlying array
		logIntSlice("result slice of appending", &s5)
		logIntSlice("original to-slice", &s4)

		s4[0] = 5	// the sharing is verified
		logIntSlice("concatenated slice after modifying original to-slice by [] operator", &s5)
		logIntSlice("original to-slice", &s4)

		log.Println("")
	}()

	log.Println("")
}

func iterateSlice() {
	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| iteration by for-range")
		log.Println("-------------------------------------------------------------------------")

		slice := []int{10, 20, 30, 40}
		logIntSlice("the slice to iterate:", &slice)

		for i, v := range slice {
			log.Printf("\tindex: %d -> value: %d", i, v)
		}

		log.Println("")
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| range provides a copy of each element")
		log.Println("-------------------------------------------------------------------------")

		slice := []int{10, 20, 30, 40}
		logIntSlice("the slice to iterate:", &slice)

		for i, v := range slice {
			log.Printf("\tindex: %d -> *value: %p, *elem: %p", i, &v, &slice[i])
		}

		log.Println("")
	}()

	func() {
		log.Println("-------------------------------------------------------------------------")
		log.Println("| iteration by traditional for-index")
		log.Println("-------------------------------------------------------------------------")

		slice := []int{10, 20, 30, 40}
		logIntSlice("the slice to iterate:", &slice)

		for i := 2; i < len(slice); i++ {
			log.Printf("\tindex: %d -> value: %d", i, slice[i])
		}

		log.Println("")
	}()

	log.Println("")
}

func main() {
	createAndInitSlice()
	manipulateSlice()
	iterateSlice()
}
