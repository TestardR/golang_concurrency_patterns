package errorhandling

// errors should be considered first-class citizens when constructing values to return
// from goroutines. If your goroutine can produce errors, those errors should be tightly
// coupled with your result type, and passed along through the same lines of communi‐
// cation—just like regular synchronous functions.
