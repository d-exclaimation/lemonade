//
//  gqlgen_generator.go
//  generator
//
//  Created by d-exclaimation on 1:20 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package generator

import (
	"github.com/d-exclaimation/lemonade/cli"
	"log"
	"os"
)

const (
	basicSchema = "type Query {\n" +
		"hello: String!\n" +
		"}"

	basicYml = "# Where are all the schema files located? globs are supported eg  src/**/*.graphqls\n" +
		"schema:\n" +
		"  - gql/*.graphqls\n" +
		"  - gql/*.graphql\n" +
		"\n" +
		"# Where should the generated server code go?\n" +
		"exec:\n" +
		"  filename: gql/gqlcore/generated.go\n" +
		"  package: gqlcore\n" +
		"\n" +
		"# Where should any generated models go?\n" +
		"model:\n" +
		"  filename: gql/model/models_gen.go\n" +
		"  package: model\n" +
		"\n" +
		"# Where should the resolver implementations go?\n" +
		"resolver:\n" +
		"  layout: follow-schema\n" +
		"  dir: gql\n" +
		"  package: gql\n" +
		"\n" +
		"# Optional: turn on use gqlgen tags in your models\n" +
		"# struct_tag: json\n" +
		"\n" +
		"# Optional: turn on to use []Thing instead of []*Thing\n" +
		"# omit_slice_element_pointers: false\n" +
		"\n" +
		"# Optional: set to speed up generation time by not performing a final validation pass.\n" +
		"# skip_validation: true\n" +
		"\n" +
		"# gqlgen will search for any type names in the schema in these go packages\n" +
		"# if they match it will use them, otherwise it will generate them.\n" +
		"autobind:\n" +
		"  - \"github.com/d-exclaimation/$name/gql/model\"\n" +
		"\n" +
		"# This section declares type mapping between the GraphQL and go type systems\n" +
		"#\n" +
		"# The first line in each type will be used as defaults for resolver arguments and\n" +
		"# modelgen, the others will be allowed when binding to fields. Configure them to\n" +
		"# your liking\n" +
		"models:\n" +
		"  ID:\n" +
		"    model:\n" +
		"      - github.com/99designs/gqlgen/graphql.ID\n" +
		"      - github.com/99designs/gqlgen/graphql.Int64\n" +
		"      - github.com/99designs/gqlgen/graphql.IntID\n" +
		"  Int:\n" +
		"    model:\n" +
		"      - github.com/99designs/gqlgen/graphql.Int\n" +
		"      - github.com/99designs/gqlgen/graphql.Int64\n" +
		"      - github.com/99designs/gqlgen/graphql.Int32"

	initModel = "package model"
)

func GqlgenTemplate(name string) {
	GoTemplate(name)
	err := os.Mkdir("./"+name+"/graphql", 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = os.Mkdir("./"+name+"/graphql/core", 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = os.Mkdir("./"+name+"/graphql/model", 0755)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Write("./"+name+"/graphql/model/model_gen.go", initModel)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Write("./"+name+"/graphql/schema.graphql", basicSchema)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	err = cli.Write("./"+name+"/gqlgen.yml", basicYml)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
}
