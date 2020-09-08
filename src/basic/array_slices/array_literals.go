package main
import "fmt"

func main () {
	//var arrayKeyValue = [5]string{3: "Chris", 4: "Ron"}
	var arrayKeyValue = []string{3: "Chris", 4: "Ron"}

	// var arrayAge = [5]int{18, 20, 15, 22, 16}
	// var arrayLazy = [...]int{5, 6, 7, 8, 22}
	// var arrayLazy = []int{5, 6, 7, 8, 22}

	

	for i := 0; i < len(arrayKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrayKeyValue[i])
	}

}