package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

const (
	defaultTTL = -time.Hour * 24 * 30
)

var (
	paths []string
	ttl   time.Duration
)

func main() {
	// cobra command
	rootCmd := &cobra.Command{
		Use:   "go-cleanup",
		Short: "go-cleanup is an utility to clean up files after certain time",
		Run:   cmdHandler,
	}

	rootCmd.Flags().StringArrayVarP(&paths, "paths", "p", paths, "paths to process")
	err := rootCmd.MarkFlagRequired("paths")
	if err != nil {
		panic(err)
	}
	rootCmd.Flags().DurationVarP(&ttl, "ttl", "t", defaultTTL, "expire time")

	err = rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func cmdHandler(cmd *cobra.Command, args []string) {
	for _, path := range paths {
		err := filepath.Walk(path, walkHandler)
		if err != nil {
			panic(err)
		}
	}
}

func walkHandler(path string, file os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	var baseline time.Time
	if ttl == defaultTTL {
		baseline = time.Now().AddDate(0, -1, 0)
	} else {
		baseline = time.Now().Add(ttl)
	}

	if file.ModTime().Before(baseline) {
		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	return nil
}
