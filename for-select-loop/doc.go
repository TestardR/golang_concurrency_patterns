package forselectloop

func main() {
	for { // Either loop infinitely or range over something
		select {
		// Do some work with channels
		}
	}

	// Sending iteration variables out on a channel
	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		}
	}

	// Looping infinitely waiting to be stopped
	// If the done channel isn’t closed, we’ll exit the select statement and continue on
	// to the rest of our for loop’s body.
	for {
		select {
		case <-done:
			return
		default:
		}
		// Do non-preemptable work
	}

	// The second variation embeds the work in a default clause of the select statement
	// When we enter the select statement, if the done channel hasn’t been closed, we’ll
	// execute the default clause instead.
	for {
		select {
		case <-done:
			return
		default:
			// Do non-preemptable work
		}
	}

}
