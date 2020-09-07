package usecases

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/GomuGomuMan/go-cleanup/internal/config"
)

const (
	defaultTTL = -time.Hour * 24 * 30
)

type Directory struct {
	Path config.Path     `json:"path"`
	TTL  config.Duration `json:"ttl,omitempty"`
}

func (d Directory) Clean() error {
	var baseline time.Time
	if d.TTL.Duration == 0 {
		baseline = time.Now().Add(defaultTTL)
	} else {
		baseline = time.Now().Add(d.TTL.Duration)
	}

	err := filepath.Walk(string(d.Path), getWalkHandler(baseline, string(d.Path)))
	if err != nil {
		return fmt.Errorf("failed to walk path: %w", err)
	}

	return nil
}

func getWalkHandler(baseline time.Time, rootPath string) filepath.WalkFunc {
	return func(path string, file os.FileInfo, err error) error {
		if path == rootPath {
			return nil
		}

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
