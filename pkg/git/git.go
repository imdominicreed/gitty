package git

import (
	"github.com/go-git/go-git/v5"
)

const DefaultMain = "mainmainRef"

type Repo struct {
	*git.Repository
}

func NewRepo(path string) (*Repo, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	return &Repo{Repository: r}, nil
}
