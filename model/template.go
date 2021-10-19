//
//  template.go
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

var choicesMap = map[string][]list.Item{
	"scala": {
		obj{title: "realtime-graphql", desc: "Ahql & OverLayer realtime GraphQL API"},
		obj{title: "restful-akka", desc: "Akka HTTP Versioned RESTful API"},
		obj{title: "scala", desc: "Stable Scala 2.13 Starter project"},
		obj{title: "dotty", desc: "Bleeding edge Scala 3 Starter project"},
	},
	"go": {
		obj{title: "gqlgen", desc: "Gqlgen GraphQL API"},
		obj{title: "go", desc: "Regular go template"},
	},
	"node/typescript": {
		obj{title: "react", desc: "Single Page React app"},
		obj{title: "nextjs", desc: "Server Side Rendered React app"},
		obj{title: "node", desc: "Node js Server"},
	},
	"elixir": {
		obj{title: "phoenix", desc: "Phoenix MVC RESTful API"},
	},
}

type Templates struct {
	list    list.Model
	choices []list.Item
	project *ProjectSetup
}

func NewTemplate(project *ProjectSetup) Templates {
	choices := choicesMap[project.Lang]
	l := NewModel(choices)
	l.Title = "Project Template"
	l.Styles.Title = titleStyle
	return Templates{
		list:    l,
		choices: choices,
		project: project,
	}
}

func (g Templates) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (g Templates) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			g.project.SetTemplate(g.choices[g.list.Cursor()].FilterValue())
			return g, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	var cmd tea.Cmd
	g.list, cmd = g.list.Update(msg)
	return g, cmd
}

func (g Templates) View() string {
	return listStyle.Render(g.list.View())
}
