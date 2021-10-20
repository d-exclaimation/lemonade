//
//  textfield.view.go
//  model
//
//  Created by d-exclaimation on 11:39 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

// TextFieldView
//
// CLI TextFieldView
type TextFieldView struct {
	textInput textinput.Model // Text input from the Bubble-tea CLI
	err       error           // Saved error
	project   *ProjectSetup   // Project setup to be modified
}

// NewTextFieldView
//
// Create a new TextFieldView
func NewTextFieldView(project *ProjectSetup) TextFieldView {
	ti := textinput.NewModel()
	ti.Placeholder = "example-app"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return TextFieldView{
		textInput: ti,
		err:       nil,
		project:   project,
	}
}

// Init
//
// Initial state for the bubble-tea cli
func (t TextFieldView) Init() tea.Cmd {
	return textinput.Blink
}

// Update
//
// Render update for bubble-tea
func (t TextFieldView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, 51, 67:
			os.Exit(1)
			return t, tea.Quit

		case tea.KeyEnter, tea.KeySpace:
			t.project.Name = t.textInput.Value()
			return t, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		t.err = msg
		return t, nil
	}

	t.textInput, cmd = t.textInput.Update(msg)
	return t, cmd
}

// View
//
// Render TextFieldView as string
func (t TextFieldView) View() string {
	return fmt.Sprintf(
		"What’s name of the project?\n\n%s\n\n%s",
		t.textInput.View(),
		"(q to quit)",
	) + "\n"
}
