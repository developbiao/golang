package main
import "fmt"

func main() {
	array := [3]float64{7.0, 3.14, 4.88}
	x := Sum(&array)
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is :%f", x)

}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return sum
}