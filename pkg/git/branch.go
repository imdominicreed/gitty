package git

import (

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type RootCommit struct {
  branch []*plumbing.Reference
  *object.Commit
}
func (r *Repo) LoadBranches() error {
  branches, err := r.repo.Branches()
  if err != nil {
    return err 
  }
  return branches.ForEach(r.ParseBranch)
}


func (r *Repo)ParseBranch(ref *plumbing.Reference) error {
  log, err := r.repo.Log(&git.LogOptions{From: ref.Hash()})
  if err != nil {
    return err
  }


  log.ForEach(
    func(c *object.Commit) error {
      return nil
    },
  )
  return nil
}
