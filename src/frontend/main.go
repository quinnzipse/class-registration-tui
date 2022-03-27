package main

import (
  "fmt"
  "strings"
  "log"
  "os"

  "github.com/charmbracelet/lipgloss"
  tea "github.com/charmbracelet/bubbletea"
)

const (
	// In real life situations we'd adjust the document to fit the width we've
	// detected. In the case of this example we're hardcoding the width, and
	// later using the detected width only to truncate in order to avoid jaggy
	// wrapping.
	width = 96

	columnWidth = 30
)

// Style definitions.
var (

	// General.

	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	url = lipgloss.NewStyle().Foreground(special).Render

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

	activeTab = tab.Copy().Border(activeTabBorder, true)

	tabGap = tab.Copy().
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)

	// Title.

	titleStyle = lipgloss.NewStyle().
			MarginLeft(1).
			MarginRight(5).
			Padding(0, 1).
			Italic(true).
			Foreground(lipgloss.Color("#FFF7DB")).
			SetString("Lip Gloss")

	descStyle = lipgloss.NewStyle().MarginTop(1)

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

  docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

func main(){
	doc := strings.Builder{}

  // Tabs
	{
		row := lipgloss.JoinHorizontal(
			lipgloss.Top,
			activeTab.Render("All Classes"),
			tab.Render("Registered Classes"),
		)
		gap := tabGap.Render(strings.Repeat(" ", max(0, width-lipgloss.Width(row)-2)))
		row = lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
		doc.WriteString(row + "\n\n")
	}

  fmt.Println(docStyle.Render(doc.String()))

  // Create a new bubbletea program.
  p := tea.NewProgram(initialModel())
  if err := p.Start(); err != nil {
    log.Fatal(err)
    fmt.Println("Bummer, there's been an error:", err)
		os.Exit(1)
  }
}

// Define what the model is. Aka what it's data types are.
type model struct {
  choices []course
  cursor_location int 
  selected map[int]struct{}
}

type course struct{ 
	courseName 	string
	profName 	string
	days 		[5]bool
	startTime   string
	endTime		string
}

// What should the model initally look like
func initialModel() model {
	
  return model {
    choices: []course{course{"CS120", "Mitra", [5]bool{true, false, true, false, true}, "9:55", "10:50"}, course{"CS220", "SauppeA", [5]bool{true, false, true, false, true}, "11:00", "11:55"}},
    selected: make(map[int]struct{}),
  }
}

// Define any initial IO here.
func (m model) Init() tea.Cmd {
  return nil
}

// Describe how the state can change. Takes a msg; Returns a model and cmd
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){ 

  // TODO: LEFT HERE https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go

  // If a key was pressed...
  switch msg := msg.(type) {
  case tea.KeyMsg:
    
    switch msg.String() {
    
    case "ctrl+c", "q":
      return m, tea.Quit

    case "up", "k":
      if m.cursor_location > 0 {
        m.cursor_location--
      }
    
    case "down", "j":
      if m.cursor_location < len(m.choices)-1 {
        m.cursor_location++
      }

    case "enter", " ":
      _, ok := m.selected[m.cursor_location]
      if ok {
        delete(m.selected, m.cursor_location)
      } else {
        m.selected[m.cursor_location] = struct{}{}
      }

    }
  }

  // return the MODIFIED model and nil (nop)
  return m, nil
}

func (m model) View() string {
  s := "Welcome to the Class Registry!\n\n"

  for i, choice := range m.choices {
    cursor := " "
    if m.cursor_location == i {
      cursor = ">"
    }

    checked := " "
    if _, ok := m.selected[i]; ok {
      checked = "✓"
    }

    // "render" the row
    s += fmt.Sprintf("%s [%s] %s %s\n", cursor, checked, choice.courseName, choice.profName);

  }

  s += "\nPress q to quit.\n"

  return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}