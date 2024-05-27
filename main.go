package main

import (
	"fmt"
	"gitty/pkg/git"
	"gitty/pkg/model"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const debug = false 

func main() {
	if debug {
		rundebug()
		return
	}
	repo, err := git.NewRepo(".")
	m, err := model.NewModel(repo)
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
}

func CheckIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
