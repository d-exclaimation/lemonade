//
//  project.module.go
//  model
//
//  Created by d-exclaimation on 9:35 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"github.com/d-exclaimation/lemonade/cli"
	"github.com/d-exclaimation/lemonade/generator"
	"github.com/d-exclaimation/lemonade/utils"
	"strings"
)

// ProjectSetup
//
// State over multiple Bubble-tea CLI
type ProjectSetup struct {
	Lang     string
	Template string
	Name     string
}

// NewProject
//
// Construct a Project
func NewProject() *ProjectSetup {
	return &ProjectSetup{
		Lang:     "scala",
		Template: "realtime-graphql",
		Name:     "example-app",
	}
}

// Execute a `bin/sh` command and render a loading for it
func (p *ProjectSetup) loading(s string, t ...string) {
	cli.Tea(NewLoadingView(func() {
		cli.Run(s, t...)
	}))
}

// Execute a `sbt` Giter8 command
func (p *ProjectSetup) sbt(t ...string) {
	p.loading("sbt", append([]string{"new"}, t...)...)
}

// Setup
//
// Set up the finished project
func (p *ProjectSetup) Setup() {
	switch p.Lang {
	case "scala":
		switch p.Template {
		case "realtime-graphql":
			p.sbt("d-exclaimation/realtime-graphql.g8", "--name="+p.Name)
		case "restful-akka":
			p.sbt("akka/akka-http-quickstart-scala.g8", "--name="+p.Name)
		case "scala":
			p.sbt("scala/scala-seed.g8", "--name="+p.Name)
		case "dotty":
			p.sbt("scala/scala3.g8", "--name="+p.Name)
		}

	case "go":
		switch p.Template {
		case "gqlgen":
			cli.Tea(NewLoadingView(func() {
				generator.GqlgenTemplate(p.Name)
			}))
		case "go":
			cli.Tea(NewLoadingView(func() {
				generator.GoTemplate(p.Name)
			}))
		}

	case "node/typescript":
		switch p.Template {
		case "react":
			p.loading("yarn", "create", "react-app", p.Name, "--template typescript")
		case "nextjs":
			p.loading("yarn", "create", "next-app", "./"+p.Name, "-e with-typescript")
		case "node":
			cli.Tea(NewLoadingView(func() {
				generator.NodeGenerator(p.Name)
			}))
		}
	case "elixir":
		switch p.Template {
		case "phoenix":
			snakeCase := utils.SnakeCase(strings.Split(p.Name, "-"))
			p.loading("mix", "phx.new", snakeCase, "--no-webpack", "--no-html", "--no-gettext", "--binary-id")
		}

	default:
		cli.Tea(NewLoadingView(func() {
			cli.Run("echo 'Done!!'")
		}))
	}
}
