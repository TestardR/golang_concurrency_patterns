package confinement

// Confinement is the simple yet powerful idea of ensuring information is only ever
// available from one concurrent process.
// There are two kinds of confinement possible: ad hoc and lexical.

// Ad hoc confinement is when you achieve confinement through a convention.
// Lexical confinement involves using lexical scope to expose only the correct data and
// concurrency primitives for multiple concurrent processes to use.
