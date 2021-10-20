//
//  constant.go
//  model
//
//  Created by d-exclaimation on 3:30 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	padding  = 2
	maxWidth = 80
)

var (
	listStyle  = lipgloss.NewStyle().Padding(1, 2)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#BC24C9")).
			Padding(0, 1)
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
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
