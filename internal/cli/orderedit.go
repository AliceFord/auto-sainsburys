package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AliceFord/auto-sainsburys/internal/plan"
	tea "github.com/charmbracelet/bubbletea"
)

type orderEditor struct {
	plan   *plan.Plan
	cursor int
	input  string
	done   bool
}

func newOrderEditor(p *plan.Plan) *orderEditor {
	return &orderEditor{
		plan: p,
	}
}

func (e orderEditor) Init() tea.Cmd {
	return nil
}

func (e orderEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return e, tea.Quit

		case "enter":
			e.done = true
			return e, tea.Quit

		case "up", "k":
			if e.cursor > 0 {
				e.cursor--
				e.input = ""
			}

		case "down", "j":
			if e.cursor < len(e.plan.Items)-1 {
				e.cursor++
				e.input = ""
			}

		case "backspace":
			if len(e.input) > 0 {
				e.input = e.input[:len(e.input)-1]

				if e.input == "" {
					e.plan.Items[e.cursor].OrderQuantity = 0
				} else {
					e.applyInput()
				}
			}

		default:
			// Accept numeric input.
			if len(msg.Runes) == 1 {
				r := msg.Runes[0]

				if r >= '0' && r <= '9' {
					e.input += string(r)
					e.applyInput()
				}
			}
		}
	}

	return e, nil
}

func (e *orderEditor) applyInput() {
	if e.input == "" {
		return
	}

	n, err := strconv.Atoi(e.input)
	if err != nil {
		return
	}

	e.plan.Items[e.cursor].OrderQuantity = n
}

func (e orderEditor) View() string {
	var b strings.Builder

	b.WriteString("\n")
	b.WriteString("Order plan\n\n")

	for i, item := range e.plan.Items {
		line := fmt.Sprintf(
			"[%d] %s (%g %s)",
			item.OrderQuantity,
			item.Name,
			item.ProductUnits,
			item.Unit,
		)

		if i == e.cursor {
			b.WriteString("> " + line + "\n")
		} else {
			b.WriteString("  " + line + "\n")
		}
	}

	b.WriteString("\n")

	if e.input != "" {
		b.WriteString(fmt.Sprintf(
			"Quantity: %s\n",
			e.input,
		))
	}

	b.WriteString(
		"↑/↓ navigate • type number to change quantity • enter confirm • q cancel",
	)

	return b.String()
}

func EditOrder(p *plan.Plan) error {
	if len(p.Items) == 0 {
		return fmt.Errorf("plan contains no items")
	}

	model := newOrderEditor(p)

	result, err := tea.NewProgram(model).Run()
	if err != nil {
		return err
	}

	editor, ok := result.(orderEditor)
	if !ok {
		return fmt.Errorf("unexpected editor result")
	}

	if !editor.done {
		return fmt.Errorf("order cancelled")
	}

	return nil
}
