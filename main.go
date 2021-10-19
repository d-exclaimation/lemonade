//
// main.go
// lemonade
//
// Created by d-exclaimation on 00:00.
//

package main

import (
	"github.com/d-exclaimation/lemonade/cli"
	"github.com/d-exclaimation/lemonade/model"
)

func main() {
	project := model.NewProject()
	cli.Tea(model.LanguageOptions(project))
	cli.Tea(model.NewTemplate(project))
	cli.Tea(model.NewTextField(project))
	project.Exec()
}
