package cmd

import (
	"fmt"
	"io"
	"os"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
)

// printStack prints the error stack to w.
func printStack(w io.Writer, err error) {
	for i, err := range errors.Stack(err) {
		if i == 0 {
			fmt.Fprintln(w, err)
		} else {
			fmt.Fprintf(w, "--- %s\n", err)
		}
		for k, v := range errors.Attributes(err) {
			fmt.Fprintf(os.Stderr, "    %s=%v\n", k, v)
		}
		if ttnErr, ok := errors.From(err); ok {
			if correlationID := ttnErr.CorrelationID(); correlationID != "" {
				fmt.Fprintf(os.Stderr, "    correlation_id=%s\n", ttnErr.CorrelationID())
			}
		}
	}
}
