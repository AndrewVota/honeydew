package button

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define a message type
type MsgSubmit struct{}

type KeyMap struct {
	Submit key.Binding
}

var DefaultKeyMap = KeyMap{
	Submit: key.NewBinding(key.WithKeys("enter")),
}

type Style struct {
	ButtonFocused lipgloss.Style
	ButtonBlurred lipgloss.Style
}

var DefaultStyle = Style{
	ButtonFocused: lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("255")).
		PaddingLeft(1).
		PaddingRight(1),
	ButtonBlurred: lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("60")).
		PaddingLeft(1).
		PaddingRight(1),
}

type Model struct {
	// General settings
	KeyMap KeyMap
	focus  bool

	// Text settings
	text string

	// Styling
	StyleButtonFocused lipgloss.Style
	StyleButtonBlurred lipgloss.Style
}

func New() Model {
	return Model{
		KeyMap: DefaultKeyMap,
		focus:  false,

		text: "Button",

		StyleButtonFocused: DefaultStyle.ButtonFocused,
		StyleButtonBlurred: DefaultStyle.ButtonBlurred,
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
	if m.focus {
		return m.StyleButtonFocused.Render(m.text)
	}

	return m.StyleButtonBlurred.Render(m.text)
}

// ---

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m *Model) Toggle() {
	m.focus = !m.focus
}
