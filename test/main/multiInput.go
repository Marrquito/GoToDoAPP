package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/charmbracelet/bubbles/cursor"
// 	"github.com/charmbracelet/bubbles/textinput"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// var (
// 	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
// 	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
// 	cursorStyle         = focusedStyle.Copy()
// 	noStyle             = lipgloss.NewStyle()
// 	helpStyle           = blurredStyle.Copy()
// 	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

// 	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
// 	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
// )

// type IO_textInput struct {
// 	input     textinput.Model
// 	input_str string
// 	output    string
// }

// type model struct {
// 	focusIndex int
// 	ios        []IO_textInput
// 	cursorMode cursor.Mode
// }

// func InitModel(m model) model {

// 	var t textinput.Model

// 	for i := range m.ios {
// 		t = textinput.New()

// 		if i == 0 {
// 			t.Focus()
// 		}

// 		t.Cursor.Style = cursorStyle
// 		t.CharLimit = 32
// 		t.Placeholder = m.ios[i].input_str

// 		t.PromptStyle = focusedStyle
// 		t.TextStyle = focusedStyle

// 		if strings.Compare(m.ios[i].input_str, "Senha") == 0 || strings.Compare(m.ios[i].input_str, "Password") == 0 {
// 			t.EchoMode = textinput.EchoPassword
// 			t.EchoCharacter = '•'
// 		}

// 		m.ios[i].input = t
// 	}

// 	return m
// }

// func (m model) Init() tea.Cmd {
// 	return textinput.Blink
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "esc":
// 			return m, tea.Quit

// 		// Change cursor mode
// 		case "ctrl+r":
// 			m.cursorMode++
// 			if m.cursorMode > cursor.CursorHide {
// 				m.cursorMode = cursor.CursorBlink
// 			}
// 			cmds := make([]tea.Cmd, len(m.ios))
// 			for i := range m.ios {
// 				cmds[i] = m.ios[i].input.Cursor.SetMode(m.cursorMode)
// 			}
// 			return m, tea.Batch(cmds...)

// 		// Set focus to next input
// 		case "tab", "shift+tab", "enter", "up", "down":
// 			s := msg.String()

// 			// Did the user press enter while the submit button was focused?
// 			// If so, exit.
// 			if s == "enter" && m.focusIndex == len(m.ios) {
// 				for i := range m.ios {
// 					m.ios[i].output = m.ios[i].input.Value()
// 				}
// 				return m, tea.Quit
// 			}

// 			// Cycle indexes
// 			if s == "up" || s == "shift+tab" {
// 				m.focusIndex--
// 			} else {
// 				m.focusIndex++
// 			}

// 			if m.focusIndex > len(m.ios) {
// 				m.focusIndex = 0
// 			} else if m.focusIndex < 0 {
// 				m.focusIndex = len(m.ios)
// 			}

// 			cmds := make([]tea.Cmd, len(m.ios))
// 			for i := 0; i <= len(m.ios)-1; i++ {
// 				if i == m.focusIndex {
// 					// Set focused state
// 					cmds[i] = m.ios[i].input.Focus()
// 					m.ios[i].input.PromptStyle = focusedStyle
// 					m.ios[i].input.TextStyle = focusedStyle
// 					continue
// 				}
// 				// Remove focused state
// 				m.ios[i].input.Blur()
// 				m.ios[i].input.PromptStyle = noStyle
// 				m.ios[i].input.TextStyle = noStyle
// 			}

// 			return m, tea.Batch(cmds...)
// 		}
// 	}

// 	// Handle character input and blinking
// 	cmd := m.updateInputs(msg)

// 	return m, cmd
// }

// func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
// 	cmds := make([]tea.Cmd, len(m.ios))

// 	// Only text inputs with Focus() set will respond, so it's safe to simply
// 	// update all of them here without any further logic.
// 	for i := range m.ios {
// 		m.ios[i].input, cmds[i] = m.ios[i].input.Update(msg)
// 	}

// 	return tea.Batch(cmds...)
// }

// func (m model) View() string {
// 	var b strings.Builder

// 	for i := range m.ios {
// 		b.WriteString(m.ios[i].input.View())
// 		if i < len(m.ios)-1 {
// 			b.WriteRune('\n')
// 		}
// 	}

// 	button := &blurredButton
// 	if m.focusIndex == len(m.ios) {
// 		button = &focusedButton
// 	}
// 	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

// 	b.WriteString(helpStyle.Render("cursor mode is "))
// 	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
// 	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

// 	return b.String()
// }

// func main() {
// 	m := model{
// 		ios: []IO_textInput{
// 			{
// 				input_str: "Login (email)",
// 			},
// 			{
// 				input_str: "Senha",
// 			},
// 		},
// 	}

// 	tprogram := tea.NewProgram(InitModel(m))
// 	finalModel, err := tprogram.Run();
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1)
// 	}

// 	// Type assertion to access the final model
// 	if finalModel, ok := finalModel.(model); ok {
// 		fmt.Println("Login:", finalModel.ios[0].output)
// 		fmt.Println("Senha:", finalModel.ios[1].output)
// 	}
// }