package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

const DefaultMain = "mainmainRef"

type Repo struct {
  repo *git.Repository   
  mainLog map[string]*object.Commit
  mainRef *plumbing.Reference
}


func NewRepo(path string) (*Repo, error) {
  r, err := git.PlainOpen(path)
  if err != nil {
    return nil, err
  }
  mainBranch, err := r.Branch(DefaultMain)
  if err != nil {
    return nil, err
  }
  ref, err := r.Reference(mainBranch.Merge, true)
  if err != nil {
    return nil, err
  }

  log, err := r.Log(&git.LogOptions{From: ref.Hash()})
  if err != nil {
    return nil, err
  }
  commits := make(map[string]*object.Commit)
  log.ForEach(func(c *object.Commit) error {
    commits[c.Hash.String()] = c
    return nil
  })

  
  return &Repo{repo: r}, nil
}
