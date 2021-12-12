package confinement

import "fmt"

// We can see that the data slice of integers is available from both the loopData function
// and the loop over the handleData channel; however, by convention we’re only access‐
// ing it from the loopData function.

func main() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
