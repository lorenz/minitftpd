package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pin/tftp"
)

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	fname := filepath.Join(".", filepath.Clean(filename))
	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	log.Printf("Sending %v", fname)
	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return err
	}
	fmt.Printf("%d bytes sent\n", n)
	return nil
}

// writeHandler is called when client starts file upload to server
func writeHandler(filename string, wt io.WriterTo) error {
	log.Printf("Rejecting write to %v", filename)
	return errors.New("writing is unsupported")
}

func main() {
	s := tftp.NewServer(readHandler, writeHandler)
	s.SetTimeout(10 * time.Second)
	err := s.ListenAndServe(":69")
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}

}
