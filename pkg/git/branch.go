package git

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Branch struct {
  Reference *plumbing.Reference
  Commits []*object.Commit
}

func (r *Repo) LoadBranches() ([]Branch, error) {
  branchIter, err := r.Branches()
  if err != nil {
    return nil, err 
  }


  branches := []Branch{}
  branchIter.ForEach(func(ref *plumbing.Reference) error {
    branch, err := r.ParseBranch(ref)
    if err != nil {
      return err
    }
    branches = append(branches, branch)
    return nil
  })
  return branches, nil
}


func (r *Repo) ParseBranch(ref *plumbing.Reference) (Branch, error) {
  log, err := r.Log(&git.LogOptions{From: ref.Hash()})
  if err != nil {
    return Branch{}, err
  }

  b := Branch{Reference: ref, Commits: []*object.Commit{}} 

  log.ForEach(
    func(c *object.Commit) error {
      b.Commits = append(b.Commits, c)
      fmt.Printf("Ref: %s commit :%s\n", ref.Name().Short(), c.Hash.String())
      return nil
    },
  )
  return b, nil 
}

type GraphCommit struct {
  *object.Commit
  BranchTips []*Branch
  Children map[string]*GraphCommit
  Parent   map[string]*GraphCommit
}

type LogGraph struct {
  RootCommits []*GraphCommit
  Commits map[string]*GraphCommit
}

func (r *Repo) BuildGraph(branches []Branch) (*LogGraph){
  graph := LogGraph{RootCommits: []*GraphCommit{}, Commits: make(map[string]*GraphCommit)}
  for _, b := range branches {
    childCommit := graph.getOrCreateCommit(b.Commits[0]) 
    childCommit.BranchTips = append(childCommit.BranchTips, &b)
    for _, branchCommit := range b.Commits[1:] {
      commit := graph.getOrCreateCommit(branchCommit)
      childCommit.AddParent(commit)
    } 
  }
  for _, commit := range graph.Commits {
    if len(commit.Parent) == 0 {
      graph.RootCommits = append(graph.RootCommits, commit)
    }
  }
  
  return &graph
}

func (l *LogGraph) getOrCreateCommit(commit *object.Commit) *GraphCommit {
  graphCommit, ok := l.Commits[commit.Hash.String()]
  if !ok { 
    graphCommit = newGrapCommit(commit)
    l.Commits[commit.Hash.String()] = graphCommit
  }
  return graphCommit
}

func (l *LogGraph) String() string {
  s := "Root Commits: "
  for _, commit := range l.RootCommits {
    s += fmt.Sprintf("%s ", commit.Hash.String())
  }
  return s
}

func (g *GraphCommit) AddParent(p *GraphCommit) {
  g.Parent[p.Hash.String()] = p 
  p.Children[g.Hash.String()] = g
}


func newGrapCommit(commit *object.Commit) *GraphCommit {
  return &GraphCommit{Commit: commit, Children: make(map[string]*GraphCommit), Parent: make(map[string]*GraphCommit), BranchTips: []*Branch{}}
}


