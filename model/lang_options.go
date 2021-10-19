//
//  lang_options.go.go
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

type LangOptions struct {
	list    list.Model
	choices []list.Item
	project *ProjectSetup
}

func LanguageOptions(project *ProjectSetup) LangOptions {
	choices := []list.Item{
		obj{title: "scala", desc: "Scala (akka-http & friends), best OOP + FP Language, good for most including GraphQL"},
		obj{title: "go", desc: "Go, lightweight practical language, good for container"},
		obj{title: "node/typescript", desc: "Node.js Typescript, Simple, easy to use language, good for small project"},
		obj{title: "elixir", desc: "Elixir, Concurrent focussed language, good for websocket with no computation"},
	}
	l := NewModel(choices)
	l.Title = "Programming language"
	l.Styles.Title = titleStyle
	return LangOptions{
		list:    l,
		choices: choices,
		project: project,
	}
}

func (o LangOptions) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (o LangOptions) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			o.project.SetLang(o.choices[o.list.Cursor()].FilterValue())
			return o, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	var cmd tea.Cmd
	o.list, cmd = o.list.Update(msg)
	return o, cmd
}

func (o LangOptions) View() string {
	return listStyle.Render(o.list.View())
}
