package main

import(
	"fmt"
	"strconv"
	"strings"
	"cli_weather/latlong"
	"cli_weather/forecast"
	"cli_weather/wrapper"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewport viewport.Model
	weather []string
	textInput textinput.Model
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter the State and City you want to check:"
	ti.Focus()

	vp := viewport.New(100, 20)
	vp.Style = lipgloss.NewStyle().
		Align(lipgloss.Left).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("49")).
		Foreground(lipgloss.Color("154"))

	vp.SetContent(`Welcome to cli_weather. Look up the forecast for a state and city.`)

	return model {
		textInput: ti,
		weather: []string{},
		viewport: vp,
	}
}

func(m model) Init() tea.Cmd {
	return textinput.Blink
}

func(m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var(
		tiCmd tea.Cmd
		vpCmd tea.Cmd
		state, city, cityfrm string
	)

	m.textInput, tiCmd = m.textInput.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textInput.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			if m.textInput.Value() == "exit" || m.textInput.Value() == "quit" {
				fmt.Println(m.textInput.Value())
				return m, tea.Quit
			}
			input := strings.Fields(m.textInput.Value())
			state = input[0]

			if len(input) > 2 {
				for i := 1; i < len(input); i++ {
					city += input[i]+" "
					if i < len(input)-1 {
						cityfrm += input[i]+"%20"
					} else {
						cityfrm += input[i]
					}
				}
			} else {
				cityfrm = input[1]
				city += input[1]
			}

			m.weather = nil
			m.viewport.SetContent(strings.Join(m.weather, ""))

			isEmpty := latlong.CheckIfEmpty(state, cityfrm)

			if isEmpty == 1 {
				m.weather = append(m.weather, "There appears to be an error, please check your spelling.\n")
				m.weather = append(m.weather, "\n"+state+", "+city+"\n")
				m.viewport.SetContent(strings.Join(m.weather, ""))
			} else {
				m.weather = append(m.weather, state+", "+city+"\n")
				for _, p := range forecast.GetForecast(state, cityfrm) {
					m.weather = append(m.weather, p.Name+"\n")
					m.weather = append(m.weather, "Temprature: "+strconv.Itoa(p.Temp)+"\n")
					theForecast := wrapper.WrapString(p.Forecast, 15)
					m.weather = append(m.weather, "Forecast: "+theForecast+"\n")
					m.weather = append(m.weather, "------------------------------------------\n")
					
					m.viewport.SetContent(strings.Join(m.weather, "\n"))
				}
			}

			m.textInput.Reset()
			m.viewport.GotoTop()
		}
	}

	return m, tea.Batch(tiCmd, vpCmd)

}

func(m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textInput.View(),
	) + "\n\n"
}

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}

}