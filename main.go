package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

func main() {
	// cobra command


	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	root := fmt.Sprintf("%s/Downloads", usr.HomeDir)

	err = filepath.Walk(root, walkHandler)
	if err != nil {
		panic(err)
	}
}

func walkHandler(path string, file os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	baseline := time.Now().AddDate(0, -1, 0)
	if file.ModTime().Before(baseline) {
		if err := os.Remove(path); err != nil {
			return err
		}
	}

	return nil
}
