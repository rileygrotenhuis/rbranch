package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item string

func (i item) FilterValue() string { return "" }

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		exec.Command("git", "checkout", m.choice).Run()
		return quitTextStyle.Render(fmt.Sprintln())
	}
	if m.quitting {
		return quitTextStyle.Render("Until next time...")
	}
	return "\n" + m.list.View()
}

func aggregateBranches(commandOutput []byte) []string {
	var branches []string

	lines := strings.Split(string(commandOutput), "\n")

	for _, line := range lines {
		branch := strings.TrimSpace(strings.TrimPrefix(line, "*"))

		if len(branch) > 0 {
			branches = append(branches, branch)
		}
	}

	return branches
}

func buildSelectionListItems(branches []string) []list.Item {
	var items []list.Item

	for _, branch := range branches {
		items = append(
			items,
			list.Item(item(branch)),
		)
	}

	return items
}

func main() {
	output, err := exec.Command("git", "branch").CombinedOutput()

	if err != nil {
		fmt.Println("fatal: not a git repository (or any of the parent directories): .git")
		return
	}

	var branches []string = aggregateBranches(output)

	var selectionItems []list.Item = buildSelectionListItems(branches)

	selectionList := list.New(selectionItems, itemDelegate{}, 20, 15)
	selectionList.Title = "Git Branches:"

	m := model{list: selectionList}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("FATAL ERROR")
		os.Exit(1)
	}
}
