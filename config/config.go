package config

import (
	"encoding/json"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

type Path string

func (d *Path) UnmarshalJSON(b []byte) error {
	var path string
	err := json.Unmarshal(b, &path)
	if err != nil {
		return err
	}

	if strings.HasPrefix(path, "~") {
		user, err := user.Current()
		if err != nil {
			return err
		}

		path = strings.Replace(path, "~", user.HomeDir, 1)
		*d = Path(path)
		return nil
	}

	path, err = filepath.Abs(path)
	if err != nil {
		return err
	}
	*d = Path(path)

	return nil
}

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var durStr string
	err := json.Unmarshal(b, &durStr)
	if err != nil {
		return err
	}

	d.Duration, err = time.ParseDuration(durStr)
	if err != nil {
		return err
	}

	return nil
}
