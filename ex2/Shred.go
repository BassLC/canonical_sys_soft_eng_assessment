// @FIX: Change package name when moving to a library.
//
//	`package` is only `main` so we can generate a
//	binary for user testing
package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func Shred(path string) error {
	// Open file in Write mode (we only need writing permissions)
	file, err := os.OpenFile(path, os.O_RDWR, 0200)
	if err != nil {
		return err
	}

	// @NOTE: We can't defer closing the file since we are going to do
	// it by hand in the future.

	// Get Stat information, so we can get the overall file size
	fileStat, err := file.Stat()
	if err != nil {
		file.Close()
		return err
	}

	// Write random bytes to file 3 times
	for i := 0; i < 3; i++ {
		_, err := io.CopyN(file, rand.Reader, fileStat.Size())
		if err != nil {
			file.Close()
			return err
		}

		// Don't forget to reset the cursor
		file.Seek(0, 0)
	}

	// @NOTE: By default, `io` functions in Go don't buffer anything,
	// meaning that they write directly to the file.  This means we
	// don't have to worry about our changes not being reflected onto
	// the actual disk, due to some "cleverness" of the
	// language. However, we still must force the OS to flush all
	// writes.
	if err = file.Sync(); err != nil {
		return err
	}

	// Close and delete the file
	if err = file.Close(); err != nil {
		return err
	}

	if err = os.Remove(path); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: Shred <file_name>")
		os.Exit(-1)
	}
	err := Shred(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
	}
}
