package model

import (
	"gitty/pkg/git"
	"strconv"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/m1gwings/treedrawer/tree"
)

type TreeModel struct {
  repo *git.Repo
  g *git.LogGraph
  vp viewport.Model
  lines []string
}


func (t *TreeModel) buildTree() error {
  branches, err := t.repo.LoadBranches()
  if err != nil {
    return  err
  }

  t.g = t.repo.BuildGraph(branches)  
  drawer := GraphDrawer{LogGraph: t.g, Width: DefaultCanvasWidth} 
  drawer.String() 
  t.vp.SetContent("")
  
  return nil
} 

func (m *TreeModel) Init() tea.Cmd {
  return nil
}

func (t *TreeModel) Update(msg tea.Msg) (*TreeModel, tea.Cmd) {
  var cmd tea.Cmd 
  t.vp, cmd = t.vp.Update(msg)
  return t, cmd
}

func (t *TreeModel) View() string {
  return t.vp.View()
}

func printTree(g *git.LogGraph) string {
  for _, root := range g.RootCommits {
    t := tree.NewTree(tree.NodeString(strconv.FormatInt(int64(len(root.Children)), 10)))
    buildTree(root, t)
    return t.String()
  }
  return ""
}


func buildTree(commit *git.GraphCommit, t *tree.Tree) {
  for hash, child := range commit.Children {
    tChild := t.AddChild(tree.NodeString(hash))
    buildTree(child, tChild)
  }
}


