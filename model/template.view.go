//
//  template.view.go
//  model
//
//  Created by d-exclaimation on 10:44 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

// TemplatesView
//
// Multi-item list select rendered by Bubbles-tea.
type TemplatesView struct {
	list    list.Model    // Inner list component
	choices []list.Item   // Choices given
	project *ProjectSetup // Project to be modified
}

// NewTemplateView
//
// Construct a new Template
func NewTemplateView(project *ProjectSetup) TemplatesView {
	choices := choicesMap[project.Lang]
	l := NewModel(choices)
	l.Title = "Project Template"
	l.Styles.Title = titleStyle
	return TemplatesView{
		list:    l,
		choices: choices,
		project: project,
	}
}

// Init
//
// Initial state for the bubble-tea cli
func (g TemplatesView) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

// Update
//
// Render update for bubble-tea
func (g TemplatesView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		topGap, rightGap, bottomGap, leftGap := listStyle.GetPadding()
		g.list.SetSize(msg.Width-leftGap-rightGap, msg.Height-topGap-bottomGap)

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.Type {
		// These keys should exit the program.
		case tea.KeyCtrlC, 51, 67:
			os.Exit(1)
			return g, tea.Quit
		case tea.KeyEnter, tea.KeySpace:
			g.project.Template = g.choices[g.list.Cursor()].FilterValue()
			return g, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	var cmd tea.Cmd
	g.list, cmd = g.list.Update(msg)
	return g, cmd
}

// View
//
// Render TextFieldView as string
func (g TemplatesView) View() string {
	return listStyle.Render(g.list.View())
}
