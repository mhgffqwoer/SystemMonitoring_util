package render

// This example demonstrates various Lip Gloss style and layout features.

import (
	"PSUtil/calculate"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	//"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/term"
)

// Определения стилей.
var (
	system_monitoring = ""
	cpuPercent = make([]float64, 8)

	// General
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	// Tabs.
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(highlight).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true)

	tabGap = tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	// Status Bar.
	treadsBarStyle = lipgloss.NewStyle().
			Margin(0, 3, 0, 0).
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	treadsNameStyle = lipgloss.NewStyle().
			Inherit(treadsBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			Align(lipgloss.Left)
			//MarginRight(1)

	treadsInfoStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Background(lipgloss.Color("#A550DF")).
			Align(lipgloss.Right).
			Width(9)

	treadsLoadStyle = lipgloss.NewStyle().
			Inherit(treadsBarStyle).
			Align(lipgloss.Left)

	// // Page.

	docStyle = lipgloss.NewStyle().Padding(0, 2, 0, 2)
)

func DownloadLine(lineSize int, info float64) string {
	x1 := []float64{0.95, 0.36, 0.36} // RGB для #F25D94
	x0 := []float64{0.76, 0.95, 0.36} // RGB для #14F9D5

	count := int(math.Round(float64(lineSize) * (0.01 * info)))
	grid := make([][]float64, count)
	for i := range grid {
		grid[i] = make([]float64, 3)
		if len(grid) == 1 {
			grid[0] = x0
			break
		}
		for j := 0; j < 3; j++ {
			grid[i][j] = x0[j] + (x1[j]-x0[j])*float64(i)/float64(count-1)
		}
	}

	loadLine := ""
	for i := 0; i < count; i++ {
		color := grid[i]
		loadLine += fmt.Sprintf("\033[38;2;%d;%d;%dm>\033[0m", int(color[0]*255), int(color[1]*255), int(color[2]*255))
	}
	return loadLine
}

func SystemMonitoring(update bool) string {
	if update || system_monitoring == "" {
		select {
			case cpuPercent = <- calculate.CpuCalculateCh:
			default:
		}
		width, _, _ := term.GetSize(int(os.Stdout.Fd()))
		columnWidth := (width - 4) / 2 - 1
		doc := strings.Builder{}

		// Tabs
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			activeTab.Render("System Monitoring"),
			tab.Render("Help Menu"),
		)
		gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(1.0, row, gap)
		doc.WriteString(row + "\n\n")

		// Status bar
		treadsCount := len(cpuPercent)

		for i := 0; i < treadsCount / 2; i++ {
			w := lipgloss.Width
			bars := make([]string, 2)

			for j := 0; j < 2; j++ {
				treadName := treadsNameStyle.
					Render(fmt.Sprintf("Tread %d", i + (treadsCount / 2) * j))
				treadInfo := treadsInfoStyle.
					Render(fmt.Sprintf("%.2f%%", cpuPercent[i + (treadsCount / 2) * j]))
				diff := columnWidth - w(treadName) - w(treadInfo)
				traedLoad := treadsLoadStyle.
					Width(diff).
					Render(DownloadLine(diff, cpuPercent[i + (treadsCount / 2) * j]))
				bars[j] = lipgloss.JoinHorizontal(lipgloss.Top,
					treadName,
					traedLoad,
					treadInfo,
				)
				bars[j] = treadsBarStyle.Width(columnWidth).Render(bars[j])
			}

			doc.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, bars[0], bars[1]) + "\n\n")
		}

		if width > 0 {
			docStyle = docStyle.MaxWidth(width)
		}

		system_monitoring = fmt.Sprint(docStyle.Render(doc.String() + "\nHello World!!"))
	}
	return system_monitoring
}