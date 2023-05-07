package textinput

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dop251/goja"
	. "github.com/invowk/invowk-cli/internal/type"
	"log"
)

var vm = goja.New()

func Bubble() string {
	p := tea.NewProgram(initialModel())

	m, err := p.StartReturningModel()

	if err != nil {
		log.Fatal(err)
	}

	if m, ok := m.(Model); ok && m.value != "" {
		return m.value
	}

	return ""
}

type errMsg error

type Model struct {
	value     string
	textInput textinput.Model
	err       error
	isValid   bool
}

type MyTest string

type CreatorModel struct {
	PromptText           *PromptText
	HelpText             *HelpText
	PlaceholderValue     *PlaceholderValue
	DefaultValue         *DefaultValue
	ValidationRegex      *ValidationRegex // mutually exclusive with ValidationJs
	ValidationJs         *ValidationJs    // mutually exclusive with ValidationRegex
	ValidationMinSize    *ValidationMinSize
	ValidationMaxSize    *ValidationMinSize
	ValidationCanBeBlank *ValidationCanBeBlank
	ValidationIsRequired *ValidationIsRequired
}

func initialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		textInput: ti,
		err:       nil,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.value = m.textInput.Value()
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	//m.value = m.textInput.Value()
	return m, cmd
}

func (m Model) View() string {
	const SCRIPT = `
function f() {
    return '(esc to' + ' quit)';
}
`
	_, err := vm.RunString(SCRIPT)
	if err != nil {
		panic(err)
	}

	var fn func() string
	err = vm.ExportTo(vm.Get("f"), &fn)
	if err != nil {
		panic(err)
	}

	if m.value == "" {
		return fmt.Sprintf(
			"What’s your favorite Pokémon?\n\n%s\n\n%s",
			m.textInput.View(),
			fn(),
		) + "\n"
	} else {
		return fmt.Sprintf(
			"What’s your favorite Pokémon?\nAnswer: %s\n\n",
			m.value,
		)
	}
}
