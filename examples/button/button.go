package main

import (
	"log"

	"github.com/andrewvota/honeydew/button"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width  int
	height int
	button button.Model
}

func New() model {
	var b = button.New()
	b.Focus()

	return model{
		width:  0,
		height: 0,
		button: b,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
			// case " ", "enter":
			// 	m.button.Toggle()
		}
	}

	cmds = append(cmds, cmd)

	m.button, cmd = m.button.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, m.button.View())
}

func main() {
	c := New()
	p := tea.NewProgram(c, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
