//
//  node.template.go
//  generator
//
//  Created by d-exclaimation on 2:01 AM.
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
	usePrisma = func() string {
		if config.NodeUsePrisma {
			return "COPY prisma ./prisma/"
		}
		return ""
	}()

	dockerNode = utils.P("",
		"FROM node:current-alpine3.11 as builder",
		"",
		"RUN mkdir /app",
		"WORKDIR /app",
		"",
		"COPY package.json ./",
		"COPY yarn.lock ./",
		usePrisma,
		"",
		"RUN yarn install --production",
		"",
		"COPY . .",
		"",
		"RUN yarn build",
		"",
		"FROM node:current-alpine3.11",
		"",
		"COPY --from=builder /app/node_modules ./node_modules",
		"COPY --from=builder /app/package*.json ./",
		"COPY --from=builder /app/dist ./dist",
		"",
		"ENV NODE_ENV production",
		"",
		"CMD [ \"yarn\", \"start\" ]",
	)
)

func index(name string) string {
	return utils.P("",
		"//",
		"// index.ts",
		"// "+name+"",
		"//",
		"// Created by "+config.GithubName+" on 00:00.",
		"//",
		"",
		"async function main() {}",
		"",
		"main();",
	)
}

func NodeGenerator(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	cli.Run("npm", "init", "-y")

	err = cli.Move("./package.json", "./"+name+"/package.json")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	err = os.Mkdir("./"+name+"/src", 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Write("./"+name+"/src/index.ts", index(name))
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Write("./"+name+"/Dockerfile", dockerNode)
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
