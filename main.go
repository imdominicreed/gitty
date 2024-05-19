package main

import (
	"fmt"
	"gitty/pkg/git"
	"gitty/pkg/model"
	"os"

	tea "github.com/charmbracelet/bubbletea"
  
)

const debug = true

func main() {
  if debug {
    rundebug()
    return
  }
  m, err := model.NewModel()
  if err != nil {
    fmt.Printf("Error running model: %s", err.Error())
    os.Exit(1)
  }
  p := tea.NewProgram(m, tea.WithAltScreen())
  if _, err := p.Run(); err != nil {
    fmt.Printf("Error running model: %s", err.Error())
    os.Exit(1)
  }
}

func rundebug() {
  repo, err := git.NewRepo(".")
  CheckIfErr(err)
  
  repo.LoadTree()
}


func CheckIfErr(err error) {
  if err != nil {
    panic(err)
  }
}
