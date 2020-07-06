package usecases

import "errors"

type Settings struct {
	Directories  []Directory  `json:"directories,omitempty"`
	Repositories []Repository `json:"repositories,omitempty"`
}

func (s Settings) Validate() error {
	if len(s.Directories) == 0 && len(s.Repositories) == 0 {
		return errors.New("no directories or repositories are specified")
	}

	return nil
}
