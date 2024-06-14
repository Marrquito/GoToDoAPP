package textInput

import (
	"fmt"
	"log"
	"regexp"

	"github.com/Marrquito/GoToDoAPP/cmd/utils"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().Background(lipgloss.Color("#ff1cca")).Foreground(lipgloss.Color("#ffb5ff")).Bold(true).Padding(0, 1, 0)
	errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8700")).Bold(true).Padding(0, 0, 0)
)

func main() {
	var output utils.Output_str
	p := tea.NewProgram(InitModel(&output, "Meu Header aqui"), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	errMsg error
)

type model struct {
	header    string
	textInput textinput.Model
	output    *utils.Output_str
	err       error
}

// sanitizeInput verifies that an input text string gets validated
func sanitizeInput(input string) error {
	matched, err := regexp.Match("^[a-zA-Z0-9_-]+$", []byte(input))
	if !matched {
		return fmt.Errorf("string violates the input regex pattern, err: %v", err)
	}
	return nil
}

func InitModel(output *utils.Output_str, header string) model {
	ti := textinput.New()
	ti.Placeholder = "Your text here"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Validate = sanitizeInput

	return model{
		textInput: ti,
		output:    output,
		header:    titleStyle.Render(header),
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.output.Update(m.textInput.Value())
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.header,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
