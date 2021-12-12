package confinement

import (
	"bytes"
	"fmt"
	"sync"
)

// You can see that because printData doesn’t close around the data slice,
// it cannot access it, and needs to take in a slice of byte to operate on.
// Thanks of the lexical scope, we’ve made it impossible to do the wrong thing,
// and so we don’t need to synchronize memory access or share data through communication.

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)

	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
