//
//  swift.template.go.go
//  generator
//
//  Created by d-exclaimation on 1:21 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package generator

import (
	"fmt"
	"github.com/d-exclaimation/lemonade/cli"
	"github.com/d-exclaimation/lemonade/utils"
	"log"
	"os"
)

func dockerSwift(name string) string {
	return utils.P("",
		"FROM swift:latest as builder",
		"WORKDIR /root",
		"COPY . .",
		"RUN swift build -c release",
		"",
		"FROM swift:slim",
		"COPY --from=builder /root/.build build",
		fmt.Sprintf("CMD [\"build/release/%s\"]", name),
	)
}

func SwiftGenerator(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	cli.RunUnder(name, "swift", "package", "init", "--type", "executable")

	err = cli.Write("./"+name+"/Dockerfile", dockerSwift(name))
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
