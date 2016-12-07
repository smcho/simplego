package main

import (
	"log"
	"os"
	"strconv"
	"bytes"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lmicroseconds)
}

func declAndInitArray() {
	func() {
		var array [5]int
		log.Printf("array, just declared: %v", array)

		var parray [5]*int
		log.Printf("array of pointers, just declared: %v", parray)
	}()

	func() {
		array := [5]int{10, 20, 30, 40, 50}
		log.Printf("array, initialized by an array literal: %v", array)
	}()

	func() {
		array := [...]int{10, 20, 30, 40, 50, 60}
		log.Printf("capacity of array can be determined by the number of values of literal: %v", array)
	}()

	func() {
		array := [5]int{1: 10, 2: 20}
		log.Printf("array initialized by literal with specific values for certain indices: %v", array)
	}()

	func() {
		array := [...]int{1: 10, 9: 20}
		log.Printf("specified no capacity, max index is respected: %v", array)
	}()
}

func manipulateArray() {
	func() {
		array := [5]int{10, 20, 30, 40, 50}
		copy := array	// array is not a reference type but a value
		array[2] = 35	// modify
		log.Printf("modify element by [] operator: %v -> %v", copy, array)
	}()

	func() {
		// passing an array as function argument by converting to a slice
		pretty := func(ap []*int) string {
			str := "["
			for i, p := range ap {
				if p == nil {
					str += "nil"
				} else {
					str += strconv.Itoa(*p)
				}

				if i + 1 < len(ap) {
					str += " "
				} else {
					str += "]"
				}
			}
			return str
		}

		// passing a pointer of an array
		// bytes.Buffer!!!
		prettyPt := func(ap *[5]*int) string {
			var buf bytes.Buffer
			buf.WriteString("[")
			for i := 0; i < len(*ap); i++ {
				if (*ap)[i] == nil {
					buf.WriteString("nil")
				} else {
					buf.WriteString(strconv.Itoa(*(*ap)[i]))
				}

				if i + 1 < len(*ap) {
					buf.WriteString(" ")
				} else {
					buf.WriteString("]")
				}
			}
			return buf.String()
		}

		array := [5]*int{0: new(int), 1: new(int)}
		log.Printf("an array of integer pointers, just initialized: %v = %s = %s", array, pretty(array[0:]), prettyPt(&array))

		*array[0] = 10
		*array[1] = 20
		log.Printf("the array of integer pointers, after modifying: %v = %s = %s", array, pretty(array[0:]), prettyPt(&array))
	}()

	func() {
		var array1 [5]string
		log.Printf("first, an array of strings, just declared: %v <- %p", array1, &array1)

		array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
		log.Printf("second, another array of strings, initialized by a literal: %v <- %p", array2, &array2)

		array1 = array2	// copy the values from array2 into array1
		log.Printf("after assign(meaing element-wise copy): first, %v <- %p vs. second, %v <- %p", array1, &array1, array2, &array2)
	}()
}

func main() {
	declAndInitArray()
	manipulateArray()
}