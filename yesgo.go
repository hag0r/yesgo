package main

import "os"

// An implementation of the yes command found in Un*x systems
// It's a fast and optimized implementation that should reach the
// same throughput as GNU yes.
// Inspired by this article: https://www.reddit.com/r/unix/comments/6gxduc/how_is_gnu_yes_so_fast/
func main() {

	var txt []byte

	// Read the text to print from the (optional) first argument
	if len(os.Args) > 1 {
		txt = []byte(os.Args[1] + "\n")
	} else {
		// If no text was given, simply print an "y"
		txt = []byte("y\n")
	}

	// We'll use a buffer of 8K
	bufLen := 8 * 1024

	buf := make([]byte, bufLen)

	// Fill as much copies of our text into the buffer as possible
	used := 0
	for used < bufLen && len(txt) <= bufLen-used {
		copy(buf[used:], txt)
		used += len([]byte(txt))
	}

	// The infinite processing loop that will endlessly print the text to StdOut
	for {
		os.Stdout.Write(buf)
	}
}
