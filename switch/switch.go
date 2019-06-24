package main

import (
	"fmt"
	"runtime"
)

func getOS() (string, string) {
	var result string
	switch os := runtime.GOOS; os {
	case "darwin":
		{
			result = "OS X.\n"
		}

	case "linux":
		{
			result = "Linux.\n"
		}

	default:
		{
			// freebsd, openbsd,
			// plan9, windows...
			result = os + ".\n"
		}
	}
	return result, "Finished"
}

// auto break
func main() {
	fmt.Print("Go runs on ")
	fmt.Print(getOS())
}
