package git

import (

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Branch struct {
  Reference *plumbing.Reference
  Commit    *object.Commit
}

func (r *Repo) LoadBranches() ([]Branch, error) {
  branchIter, err := r.References()
  if err != nil {
    return nil, err 
  }


  branches := []Branch{}
  branchIter.ForEach(func(ref *plumbing.Reference) error {
    branch, err := r.GetBranch(ref)
    if err != nil {
      return err
    }
    branches = append(branches, branch)
    return nil
  })
  return branches, nil
}


func (r *Repo) GetBranch(ref *plumbing.Reference) (Branch, error) {
  commit, err := r.CommitObject(ref.Hash())
  if err != nil {
    return Branch{}, err
  }
  b := Branch{Reference: ref, Commit: commit} 
  return b, nil 
}

type GraphCommit struct {
  *object.Commit
  BranchTips []*Branch
  Children map[string]*GraphCommit
  ParentCommits []*GraphCommit
  Length int
  Head bool
}

type LogGraph struct {
  RootCommits []*GraphCommit
  Commits map[string]*GraphCommit
}

func (r *Repo) BuildGraph(branches []Branch) (*LogGraph){
  graph := LogGraph{RootCommits: []*GraphCommit{}, Commits: make(map[string]*GraphCommit)}
  for _, b := range branches {
    commit, _ := graph.getOrCreateCommit(b.Commit) 
    commit.BranchTips = append(commit.BranchTips, &b)

    commit.Length = max(commit.Length, 0)

    graph.build(commit)
  }
  for _, commit := range graph.Commits {
    if commit.NumParents() == 0 {
      graph.RootCommits = append(graph.RootCommits, commit)
    }
  }

  head, err := r.Head()
  if err == nil {
    graph.Commits[head.Hash().String()].Head = true
  }
  return &graph
}
 
func (l *LogGraph) build(commit *GraphCommit) {
      commit.Parents().ForEach(
        func(parentGitCommit *object.Commit) error {
          parent, _ := l.getOrCreateCommit(parentGitCommit)
          
          parent.Length = max(parent.Length, commit.Length + 1)
          commit.AddParent(parent)

          l.build(commit)
          return nil
        },
      )
}

func (l *LogGraph) getOrCreateCommit(commit *object.Commit) (*GraphCommit, bool) {
  graphCommit, ok := l.Commits[commit.Hash.String()]
  if !ok { 
    graphCommit = newGrapCommit(commit)
    l.Commits[commit.Hash.String()] = graphCommit
    return graphCommit, false
  }
  return graphCommit, true
}

func (g *GraphCommit) AddParent(p *GraphCommit) {
  g.ParentCommits = append(g.ParentCommits, p)
  p.Children[g.Hash.String()] = g
}


func newGrapCommit(commit *object.Commit) *GraphCommit {
  return &GraphCommit{Commit: commit, Children: make(map[string]*GraphCommit), BranchTips: []*Branch{}}
}



