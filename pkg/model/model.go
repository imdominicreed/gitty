package model

import (
	"gitty/pkg/git"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type GitModel struct {
  ta *TreeModel 
}


func NewModel(repo *git.Repo) (*GitModel, error) {
  vp := viewport.New(1000, 300)


  t := &TreeModel{vp: vp, repo: repo}
  t.buildTree()

  return &GitModel{ta:t}, nil
}

func (m *GitModel) Init() tea.Cmd {
	return nil
}

func (m *GitModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
   switch msg := msg.(type) {
    // Is it a key press?
    case tea.KeyMsg:

        // Cool, what was the actual key pressed?
        switch msg.String() {

        // These keys should exit the program.
        case "ctrl+c", "q":
            return m, tea.Quit
        }

        default:
          var cmd tea.Cmd
          _, cmd = m.ta.Update(msg)
          return m, cmd
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}


func (m *GitModel) View() string {
  return m.ta.View()
}


