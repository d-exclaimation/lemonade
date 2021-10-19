//
//  textfield.go
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

type errMsg error

type TextField struct {
	textInput textinput.Model
	err       error
	project   *ProjectSetup
}

func NewTextField(project *ProjectSetup) TextField {
	ti := textinput.NewModel()
	ti.Placeholder = "example-app"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return TextField{
		textInput: ti,
		err:       nil,
		project:   project,
	}
}

func (t TextField) Init() tea.Cmd {
	return textinput.Blink
}

func (t TextField) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, 51, 67:
			os.Exit(1)
			return t, tea.Quit
		case tea.KeyEnter, tea.KeySpace:
			t.project.SetName(t.textInput.Value())
			return t, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		t.err = msg
		return t, nil
	}

	t.textInput, cmd = t.textInput.Update(msg)
	return t, cmd
}

func (t TextField) View() string {
	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n%s",
		t.textInput.View(),
		"(q to quit)",
	) + "\n"
}
