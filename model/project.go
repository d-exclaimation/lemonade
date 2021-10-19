//
//  project.go
//  model
//
//  Created by d-exclaimation on 9:35 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/d-exclaimation/lemonade/cli"
	"github.com/d-exclaimation/lemonade/generator"
	"github.com/d-exclaimation/lemonade/utils"
	"os"
	"strings"
)

var (
	listStyle  = lipgloss.NewStyle().Padding(1, 2)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#BC24C9")).
			Padding(0, 1)
)

type ProjectSetup struct {
	Lang     string
	Template string
	Name     string
}

func NewProject() *ProjectSetup {
	return &ProjectSetup{
		Lang:     "scala",
		Template: "realtime-graphql",
		Name:     "example-app",
	}
}

func (p *ProjectSetup) SetLang(lang string) {
	p.Lang = lang
}

func (p *ProjectSetup) SetTemplate(template string) {
	p.Template = template
}

func (p *ProjectSetup) SetName(name string) {
	p.Name = name
}

func (p *ProjectSetup) Details() string {
	return fmt.Sprintf("Project: %s-obj using Template \"%s\", named \"%s\"", p.Lang, p.Template, p.Name)
}

func (p *ProjectSetup) ExecAndLoad(s string, t ...string) {
	cmd := cli.Run(s, t...)
	cli.Tea(NewLoading(func() {
		err := cmd.Wait()
		if err != nil {
			os.Exit(1)
		}
	}))
}

func (p *ProjectSetup) Exec() {
	switch p.Lang {
	case "scala":
		giter := func(t ...string) {
			p.ExecAndLoad("sbt", append([]string{"new"}, t...)...)
		}

		switch p.Template {
		case "realtime-graphql":
			giter("d-exclaimation/realtime-graphql.g8", "--name="+p.Name)
		case "restful-akka":
			giter("akka/akka-http-quickstart-scala.g8", "--name="+p.Name)
		case "scala":
			giter("scala/scala-seed.g8", "--name="+p.Name)
		case "dotty":
			giter("scala/scala3.g8", "--name="+p.Name)
		}

	case "go":
		switch p.Template {
		case "gqlgen":
			cli.Tea(NewLoading(func() {
				generator.GqlgenTemplate(p.Name)
			}))
		case "go":
			cli.Tea(NewLoading(func() {
				generator.GoTemplate(p.Name)
			}))
		}

	case "node/typescript":
		switch p.Template {
		case "react":
			p.ExecAndLoad("yarn", "create", "react-app", p.Name, "--template typescript")
		case "nextjs":
			p.ExecAndLoad("yarn", "create", "next-app", "./"+p.Name, "-e with-typescript")
		case "node":
			cli.Tea(NewLoading(func() {
				generator.NodeGenerator(p.Name)
			}))
		}
	case "elixir":
		switch p.Template {
		case "phoenix":
			snakeCase := utils.SnakeCase(strings.Split(p.Name, "-"))
			p.ExecAndLoad("mix", "phx.new", snakeCase, "--no-webpack", "--no-html", "--no-gettext", "--binary-id")
		}
	default:
		cmd := cli.Run("echo 'Done!!'")
		cli.Tea(NewLoading(func() {
			err := cmd.Wait()
			if err != nil {
				os.Exit(1)
			}
		}))
	}
}
