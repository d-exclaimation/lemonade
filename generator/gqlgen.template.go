//
//  gqlgen.template.go
//  generator
//
//  Created by d-exclaimation on 1:20 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package generator

import (
	"github.com/d-exclaimation/lemonade/cli"
	"github.com/d-exclaimation/lemonade/utils"
	"log"
	"os"
)

var (
	basicSchema = utils.P("",
		"type Query {",
		"  hello: String!",
		"}",
	)

	basicYml = utils.P("",
		"# Where are all the schema files located? globs are supported eg  src/**/*.graphqls",
		"schema:",
		"  - gql/*.graphqls",
		"  - gql/*.graphql",
		"",
		"# Where should the generated server code go?",
		"exec:",
		"  filename: gql/gqlcore/generated.go",
		"  package: gqlcore",
		"",
		"# Where should any generated models go?",
		"model:",
		"  filename: gql/model/models_gen.go",
		"  package: model",
		"",
		"# Where should the resolver implementations go?",
		"resolver:",
		"  layout: follow-schema",
		"  dir: gql",
		"  package: gql",
		"",
		"# Optional: turn on use gqlgen tags in your models",
		"# struct_tag: json",
		"",
		"# Optional: turn on to use []Thing instead of []*Thing",
		"# omit_slice_element_pointers: false",
		"",
		"# Optional: set to speed up generation time by not performing a final validation pass.",
		"# skip_validation: true",
		"",
		"# gqlgen will search for any type names in the schema in these go packages",
		"# if they match it will use them, otherwise it will generate them.",
		"autobind:",
		"  - github.com/d-exclaimation/$name/gql/model\"\n",
		"",
		"# This section declares type mapping between the GraphQL and go type systems",
		"#",
		"# The first line in each type will be used as defaults for resolver arguments and",
		"# modelgen, the others will be allowed when binding to fields. Configure them to",
		"# your liking",
		"models:",
		"  ID:",
		"    model:",
		"      - github.com/99designs/gqlgen/graphql.ID",
		"      - github.com/99designs/gqlgen/graphql.Int64",
		"      - github.com/99designs/gqlgen/graphql.IntID",
		"  Int:",
		"    model:",
		"      - github.com/99designs/gqlgen/graphql.Int",
		"      - github.com/99designs/gqlgen/graphql.Int64",
		"      - github.com/99designs/gqlgen/graphql.Int32",
	)

	initModel = utils.P("",
		"package model",
	)
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
