package main

// import (
// 	"fmt"
// 	"os"

// 	tea "github.com/charmbracelet/bubbletea"
// )

import (
	"fmt"
	"os"

	//"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
    "PSUtil/render"
    "PSUtil/calculate"
)

type model struct {
    counter int
    update bool
}

func initialModel() model {
	return model{counter: 0, update: true}
}

func (m model) Init() tea.Cmd {
    return tick()
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(2*time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    m.update = false
    switch msg := msg.(type) {

    case tea.KeyMsg: 
        switch msg.String() {
        case "h":
            m.counter = 1 
        case "ctrl+c", "q":
            return m, tea.Quit
        case "b":
            m.counter = 0
        case tea.KeyLeft.String():
            m.counter = (m.counter + 1) % 2
        case tea.KeyRight.String():
            m.counter = (m.counter + 1) % 2
        }
    case tea.MouseMsg:
        switch msg.Button {
        case tea.MouseButtonWheelUp:
            return m, nil
        case tea.MouseButtonWheelDown:
            return m, nil
        }
    case tickMsg:
        m.update = true
        return m, tick()
    }
    return m, nil
}

func (m model) View() string {
    var result string
    if m.counter == 0 {
        result = render.SystemMonitoring(m.update)
    } else if m.counter == 1 {
        result = render.HelpMenu()
    }

    return result
}

func main() {
    calculate.CpuCalculate()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseCellMotion())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
