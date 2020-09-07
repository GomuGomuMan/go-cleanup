package usecases

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/GomuGomuMan/go-cleanup/internal/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

type Repository struct {
	Path          config.Path `json:"path"`
	Pattern       string      `json:"pattern"`
	KeepRecentNum int         `json:"keep_recent_num"`
}

func (r Repository) Clean() error {
	regex := regexp.MustCompile(r.Pattern)

	path := string(r.Path)
	repo, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("failed to open repo path %s: %w", path, err)
	}

	var branches storer.ReferenceIter
	branches, err = repo.Branches()
	if err != nil {
		return fmt.Errorf("failed to get branches for repo path %s: %w", path, err)
	}

	var matchedBranches []plumbing.ReferenceName

	err = branches.ForEach(func(reference *plumbing.Reference) error {
		name := strings.TrimPrefix(reference.Name().String(), "refs/heads/")

		if regex.MatchString(name) {
			matchedBranches = append(matchedBranches, reference.Name())
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to delete branches: %w", err)
	}

	if len(matchedBranches) <= r.KeepRecentNum {
		return nil
	}

	for i := 0; i < len(matchedBranches) && len(matchedBranches)-i > r.KeepRecentNum; i++ {
		err = repo.Storer.RemoveReference(matchedBranches[i])
		if err != nil {
			return fmt.Errorf("failed to delete branch %s: %w", matchedBranches[i], err)
		}
	}
	return nil
}
