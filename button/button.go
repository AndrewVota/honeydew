package button

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	White      = "#FFFFFF"
	LightGray  = "#D3D3D3"
	MediumGray = "#A0A0A0"
	DarkGray   = "#808080"
)

type MsgSubmit struct{}
type msgPress struct{}

type KeyMap struct {
	Submit key.Binding
}

var DefaultKeyMap = KeyMap{
	Submit: key.NewBinding(key.WithKeys("enter")),
}

type Styles struct {
	ButtonFocused lipgloss.Style
	ButtonBlurred lipgloss.Style
	ButtonPressed lipgloss.Style
}

var DefaultStyles Styles = Styles{
	ButtonFocused: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(White)).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(White)).PaddingLeft(1).PaddingRight(1),
	ButtonBlurred: lipgloss.NewStyle().Bold(false).Foreground(lipgloss.Color(LightGray)).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(LightGray)).PaddingLeft(1).PaddingRight(1),
	ButtonPressed: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(DarkGray)).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(DarkGray)).PaddingLeft(1).PaddingRight(1),
}

type Model struct {
	// General settings
	KeyMap  KeyMap
	focus   bool
	pressed bool

	// Text settings
	text string

	// Styling
	FocusedStyle lipgloss.Style
	PressedStyle lipgloss.Style
	BlurredStyle lipgloss.Style
}

func New() Model {
	return Model{
		KeyMap:  DefaultKeyMap,
		focus:   false,
		pressed: false,

		text: "Button",

		FocusedStyle: DefaultStyles.ButtonFocused,
		PressedStyle: DefaultStyles.ButtonPressed,
		BlurredStyle: DefaultStyles.ButtonBlurred,
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
			m.pressed = true
			cmd := tea.Tick(time.Millisecond*200, func(t time.Time) tea.Msg {
				return msgPress{}
			})
			return m, tea.Batch(cmd, func() tea.Msg { return MsgSubmit{} })
		}

	case msgPress:
		m.pressed = false
	}

	return m, nil
}

func (m Model) View() string {
	if m.pressed {
		return m.PressedStyle.Render(m.text)
	}

	if m.focus {
		return m.FocusedStyle.Render(m.text)
	}

	return m.BlurredStyle.Render(m.text)
}

// ---

func (m *Model) Focus() {
	m.focus = true
}

func (m *Model) Blur() {
	m.focus = false
}

func (m *Model) ToggleFocus() {
	m.focus = !m.focus
}
