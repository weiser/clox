package value

import "fmt"

type Value float64

type ValueArray []Value

func PrintValue(v Value) {
	fmt.Printf("%f", v)
}
