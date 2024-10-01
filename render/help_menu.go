package render

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

func HelpMenu() string {
	width, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	// Tabs
	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tab.Render("System Monitoring"),
		activeTab.Render("Help Menu"),
	)
	gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
	row = lipgloss.JoinHorizontal(1.0, row, gap)
	doc.WriteString(row + "\n\nHelp_menu TODO!")

	if width > 0 {
		docStyle = docStyle.MaxWidth(width)
	}

	return fmt.Sprint(docStyle.Render(doc.String()))
}