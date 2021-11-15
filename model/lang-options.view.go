//
//  lang-options.view.go.go
//  model
//
//  Created by d-exclaimation on 9:11 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

// LangOptionsView
//
// A selectable list for language choice.
type LangOptionsView struct {
	list    list.Model    // Inner bubble-tea list
	choices []list.Item   // Choices for state
	project *ProjectSetup // Project setup to be mutated
}

// NewLangOptionsView
//
// Construct a new LangOptionsView
func NewLangOptionsView(project *ProjectSetup) LangOptionsView {
	choices := []list.Item{
		obj{title: "scala", desc: "OOP + FP Languages, Nice language feature on JVM but heavier than others"},
		obj{title: "go", desc: "Practical, structured, simple, fast"},
		obj{title: "node/typescript", desc: "Simple, easy, cleaner compared to Go but slower and less powerful"},
		obj{title: "elixir", desc: "Concurrent focused language, good for websocket but lack computation speed"},
		obj{title: "swift", desc: "Sophisticated, premium-feeling, newly concurrent focused language"},
	}
	l := NewModel(choices)
	l.Title = "Programming language"
	l.Styles.Title = titleStyle
	return LangOptionsView{
		list:    l,
		choices: choices,
		project: project,
	}
}

// Init
//
// Initial state for the bubble-tea cli
func (o LangOptionsView) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

// Update
//
// Render update for bubble-tea
func (o LangOptionsView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		topGap, rightGap, bottomGap, leftGap := listStyle.GetPadding()
		o.list.SetSize(msg.Width-leftGap-rightGap, msg.Height-topGap-bottomGap)

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			os.Exit(1)
			return o, tea.Quit
		case "enter", " ":
			o.project.Lang = o.choices[o.list.Cursor()].FilterValue()
			return o, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	var cmd tea.Cmd
	o.list, cmd = o.list.Update(msg)
	return o, cmd
}

// View
//
// Render TextField as string
func (o LangOptionsView) View() string {
	return listStyle.Render(o.list.View())
}
