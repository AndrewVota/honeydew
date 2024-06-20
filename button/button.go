package button

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// Define a message type
type MsgSubmit struct{}

type KeyMap struct {
	Submit key.Binding
}

var DefaultKeyMap = KeyMap{
	Submit: key.NewBinding(key.WithKeys("enter")),
}

type Model struct {
	// General settings
	KeyMap KeyMap
	focus  bool

	// Text settings
	text string
}

func New() Model {
	return Model{
		KeyMap: DefaultKeyMap,
		focus:  false,

		text: "Button",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focus {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Submit):
			return m, func() tea.Msg { return MsgSubmit{} }
		}
	}

	return m, nil
}

func (m Model) View() string {
	return m.text
}

// ---

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}
