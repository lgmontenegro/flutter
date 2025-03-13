package helpers

import (
	"fmt"
	"os"
)

func ErrHandler(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
