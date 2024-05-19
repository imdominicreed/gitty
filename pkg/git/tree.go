package git

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/object"
)

type Tree struct {}

func (r *Repo) LoadTree() error {
  trees, err := r.repo.TreeObjects()
  if err != nil {
    return err 
  }
  fmt.Println("loaded tree")
  return trees.ForEach(ParseTree)
}


func ParseTree(tr *object.Tree) error {
  for _, entry := range tr.Entries {
    fmt.Println(entry.Name)
  }
  return nil
}
