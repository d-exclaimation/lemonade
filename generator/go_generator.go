//
//  go_generator.go.go
//  generator
//
//  Created by d-exclaimation on 12:51 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package generator

import (
	"github.com/d-exclaimation/lemonade/cli"
	"log"
	"os"
)

const (
	dockergo = "FROM golang:1.17.2-alpine3.14 as compiler\n" +
		"\n" +
		"RUN mkdir /app\n" +
		"WORKDIR /app\n" +
		"\n" +
		"COPY go.mod ./\n" +
		"COPY go.sum ./\n" +
		"\n" +
		"RUN go mod download\n" +
		"\n" +
		"COPY . .\n" +
		"\n" +
		"RUN go build -o main .\n" +
		"\n" +
		"FROM golang:1.17.2-alpine3.14\n" +
		"\n" +
		"COPY --from=compiler /app/main ./app/main\n" +
		"\n" +
		"ENV GO_ENV production\n" +
		"\n" +
		"CMD [ \"/app/main\" ]"

	gitignore = "# Logs\n" +
		"logs\n" +
		"*.log\n" +
		"\n" +
		"# Env\n" +
		".env\n" +
		"\n" +
		"# OS\n" +
		".DS_Store\n" +
		"\n# Tests\n" +
		"/coverage\n" +
		"\n# IDEs\n" +
		"/.idea\n" +
		".project\n" +
		"/.vscode"

	dockerignore = "passphrase.txt\n" +
		".gitignore\n" +
		".env\n" +
		".idea/\n" +
		"logs/\n" +
		".git/\n" +
		"*.md\n" +
		".cache"
)

func header(name string) string {
	return "//\n" +
		"// main.go\n" +
		"// " + name + "\n" +
		"//\n" +
		"// Created by d-exclaimation on 00:00.\n" +
		"//\n" +
		"\n" +
		"package main"
}

func GoTemplate(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Run("go", "mod", "init", "github.com/d-exclaimation/"+name).Wait()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
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
	err = cli.Write("./"+name+"/Dockerfile", dockergo)
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
