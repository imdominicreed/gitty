package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type GitModel struct {
  viewport viewport.Model
}


func NewModel() (*GitModel, error) {
  vp := viewport.New(78, 20)

  vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

  renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(78),
	)
	if err != nil {
		return nil, err
	}
  var test string
  for i := 0; i < 1000; i++ {
    test += fmt.Sprintf("%v\n", i)
  }
  str, err := renderer.Render(test)
  if err != nil {
    return nil, err
  }

  vp.SetContent(str)

  return &GitModel{viewport: vp}, nil
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

        default:
          var cmd tea.Cmd
          m.viewport, cmd = m.viewport.Update(msg)
          return m, cmd
        }
    }

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}


func (m *GitModel) View() string {
  return m.viewport.View() 
}


