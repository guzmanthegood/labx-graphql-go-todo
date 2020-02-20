package main

import (
	"io/ioutil"
	"os"

	graphql "github.com/guzmanweb/graphql-go"
)

func loadSchema(resolverQuery interface{}, resolverMutation interface{}) *graphql.Schema  {
	file, err := os.Open("./graphql/schema.graphql")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	schemaTxt, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	schema, err := graphql.ParseSchema(string(schemaTxt), resolverQuery, resolverMutation)
	if err != nil {
		panic(err)
	}
	return schema
}