package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/marco-souza/omg/internal/llm"
)

type model struct {
	textarea textarea.Model
	err      error
}

func MakeModel() model {
	ti := textarea.New()
	ti.Placeholder = " You are Mario from Super Mario Bros, acting as an assistant..."
	ti.Focus()

	return model{
		textarea: ti,
		err:      nil,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, textarea.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// select msg type
		switch msg.Type {

		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}

		case tea.KeyEnter:
			if !m.textarea.Focused() {
				fmt.Println("Done!")

				input := m.textarea.Value()
				fmt.Println(input)

        llm.Completion(input)

				return m, tea.Quit
			}

		case tea.KeyCtrlC:
			return m, tea.Quit

		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}

		}

	case error:
		// return if msg is an error
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	avaliableHotkeys := "(ctrl+c) to quit, (esc) to save"
	if !m.textarea.Focused() {
		avaliableHotkeys = "(enter) to save"
	}
	return fmt.Sprintf(
		"Tell me what you need.\n\n%s\n\n%s",
		m.textarea.View(),
		avaliableHotkeys,
	) + "\n\n"
}
