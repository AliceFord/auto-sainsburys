package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AliceFord/auto-sainsburys/internal/plan"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	editedQuantityStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00")) // yellow
	zeroQuantityStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")) // red
	cursorStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")) // green
)

type orderEditor struct {
	plan               *plan.Plan
	originalQuantities []int
	cursor             int
	input              string
	done               bool

	// Terminal height
	height int
}

func newOrderEditor(p *plan.Plan) *orderEditor {
	originalQuantities := make([]int, len(p.Items))

	for i, item := range p.Items {
		originalQuantities[i] = item.OrderQuantity
	}

	return &orderEditor{
		plan:               p,
		originalQuantities: originalQuantities,
	}
}

func (e orderEditor) Init() tea.Cmd {
	return nil
}

func (e orderEditor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		e.height = msg.Height
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

	start, end := e.visibleRange()

	if start > 0 {
		b.WriteString("  ↑ more\n")
	}

	for i := start; i < end; i++ {
		item := e.plan.Items[i]

		line := fmt.Sprintf(
			"%s %s (%g %s)",
			e.renderQuantity(i),
			item.Name,
			item.ProductUnits,
			item.Unit,
		)

		if i == e.cursor {
			b.WriteString(cursorStyle.Render("> ") + line + "\n")
		} else {
			b.WriteString("  " + line + "\n")
		}
	}

	if end < len(e.plan.Items) {
		b.WriteString("  ↓ more\n")
	}

	b.WriteString("\n")

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

func (e orderEditor) renderQuantity(i int) string {
	quantity := e.plan.Items[i].OrderQuantity

	text := fmt.Sprintf("[%d]", quantity)

	switch quantity {
	case e.originalQuantities[i]:
		return text
	case 0:
		return zeroQuantityStyle.Render(text)
	default:
		return editedQuantityStyle.Render(text)
	}
}

func (e orderEditor) visibleRange() (int, int) {
	// 6 lines for header, footer, and padding
	available := max(e.height-6, 1)

	if available >= len(e.plan.Items) {
		return 0, len(e.plan.Items)
	}

	start := max(e.cursor-available/2, 0)

	end := start + available
	if end > len(e.plan.Items) {
		end = len(e.plan.Items)
		start = end - available
	}

	return start, end
}
