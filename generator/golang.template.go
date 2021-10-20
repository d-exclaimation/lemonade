//
//  golang.template.go.go
//  generator
//
//  Created by d-exclaimation on 12:51 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package generator

import (
	"github.com/d-exclaimation/lemonade/cli"
	"github.com/d-exclaimation/lemonade/config"
	"github.com/d-exclaimation/lemonade/utils"
	"log"
	"os"
)

var (
	dockerGo = utils.P("",
		"FROM golang:"+config.GoVersion+"-alpine3.14 as compiler",
		"",
		"RUN mkdir /app",
		"WORKDIR /app",
		"",
		"COPY go.mod ./",
		"COPY go.sum ./",
		"",
		"RUN go mod download",
		"",
		"COPY . .",
		"",
		"RUN go build -o main .",
		"",
		"FROM golang:"+config.GoVersion+"-alpine3.14",
		"",
		"COPY --from=compiler /app/main ./app/main",
		"",
		"ENV GO_ENV production",
		"",
		"CMD [ \"/app/main\" ]",
	)

	gitignore = utils.P("",
		"# Logs",
		"logs",
		"*.log",
		"",
		"# Env",
		".env",
		"",
		"# OS",
		".DS_Store",
		"",
		"# Tests",
		"/coverage",
		"",
		"# IDEs",
		"/.idea",
		".project",
		"/.vscode",
	)

	dockerignore = utils.P("",
		"passphrase.txt",
		".gitignore",
		".env",
		".idea/",
		"logs/",
		".git/",
		"*.md",
		".cache",
	)
)

func header(name string) string {
	return utils.P("",
		"//",
		"// main.go",
		"// "+name+"",
		"//",
		"// Created by "+config.GithubName+" on 00:00.",
		"//",
		"",
		"package main",
	)
}

func GoTemplate(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	cli.Run("go", "mod", "init", "github.com/"+config.GithubName+"/"+name)

	err = cli.Move("./go.mod", "./"+name+"/go.mod")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	err = cli.Write("./"+name+"/main.go", header(name))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Write("./"+name+"/Dockerfile", dockerGo)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Write("./"+name+"/.gitignore", gitignore)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Write("./"+name+"/.dockerignore", dockerignore)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
