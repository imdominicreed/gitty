package model

import (
	"gitty/pkg/git"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
var style = lipgloss.NewStyle().
    BorderStyle(lipgloss.NormalBorder()).
    BorderForeground(lipgloss.Color("63"))

type TreeModel struct {
  repo *git.Repo
  g *git.LogGraph
  vp viewport.Model
  output string
}


func (t *TreeModel) buildTree() error {
  branches, err := t.repo.LoadBranches()
  if err != nil {
    return  err
  }

  t.g = t.repo.BuildGraph(branches)  
  drawer := GraphDrawer{LogGraph: t.g, Width: DefaultCanvasWidth} 
  t.output = drawer.String()
  
  return nil
} 

func (m *TreeModel) Init() tea.Cmd {
  return nil
}

func (t *TreeModel) Update(msg tea.Msg) (*TreeModel, tea.Cmd) {
  return t, nil 
}

func (t *TreeModel) View() string {
  return style.Render(t.output)
}



