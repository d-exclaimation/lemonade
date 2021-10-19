//
//  node_generator.go
//  generator
//
//  Created by d-exclaimation on 2:01 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package generator

import (
	"github.com/d-exclaimation/lemonade/cli"
	"log"
	"os"
)

const dockerNode = "FROM node:current-alpine3.11 as builder\n" +
	"\n" +
	"RUN mkdir /app\n" +
	"WORKDIR /app\n" +
	"\n" +
	"COPY package.json ./\n" +
	"COPY yarn.lock ./\n" +
	"COPY prisma ./prisma/\n" +
	"\n" +
	"RUN yarn install --production\n" +
	"\n" +
	"COPY . .\n" +
	"\n" +
	"RUN yarn build\n" +
	"\n" +
	"FROM node:current-alpine3.11\n" +
	"\n" +
	"COPY --from=builder /app/node_modules ./node_modules\n" +
	"COPY --from=builder /app/package*.json ./\n" +
	"COPY --from=builder /app/dist ./dist\n" +
	"\n" +
	"ENV NODE_ENV production\n" +
	"\n" +
	"CMD [ \"yarn\", \"start\" ]"

func index(name string) string {
	return "//\n" +
		"// index.ts\n" +
		"// " + name + "\n" +
		"//\n" +
		"// Created by d-exclaimation on 00:00.\n" +
		"//\n" +
		"\n" +
		"async function main() {}\n" +
		"\n" +
		"main();"
}

func NodeGenerator(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Run("npm", "init", "-y").Wait()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
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
