package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

const (
	defaultTTL = -time.Hour * 24 * 30

	flagPath = "path"
)

var (
	paths           []string
	ttl             time.Duration
)

func main() {
	// cobra command
	rootCmd := &cobra.Command{
		Use:   "go-cleanup",
		Short: "go-cleanup is an utility to clean up files after certain time",
		Run:   cmdHandler,
	}

	rootCmd.Flags().StringArrayVarP(&paths, flagPath, "p", paths, "paths to process")
	err := rootCmd.MarkFlagRequired(flagPath)
	if err != nil {
		panic(err)
	}
	rootCmd.Flags().DurationVarP(&ttl, "ttl", "t", defaultTTL, "expire time")
	if int64(ttl) >= 0 {
		panic("ttl shouldn't be positive")
	}

	err = rootCmd.Execute()
	if err != nil && err != filepath.SkipDir {
		panic(err)
	}
}

func cmdHandler(cmd *cobra.Command, args []string) {
	for _, path := range paths {
		err := filepath.Walk(path, getWalkHandler())
		if err != nil {
			panic(err)
		}
	}
}

func getWalkHandler() filepath.WalkFunc {
	baseline := time.Now().Add(ttl)

	return func(path string, file os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to walk path: %w", err)
		}

		if file.ModTime().Before(baseline) {
			if err := os.RemoveAll(path); err != nil {
				return fmt.Errorf("failed to remove file/dir: %w", err)
			}
		}

		return nil
	}
}
