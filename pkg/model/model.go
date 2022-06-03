package model

import (
  // "fmt"

  // "github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/muesli/termenv"
)

type Model struct {
  //
  Text string
}

func (m Model) Init() tea.Cmd {
  return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return m, tea.Quit
}

func (m Model) View() string {
  s := "View"

  return s
}
