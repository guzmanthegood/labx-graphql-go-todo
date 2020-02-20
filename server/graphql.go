package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"labx-graphql-go-todo/graphql/resolver"
	"net/http"
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

func graphQLService(w http.ResponseWriter, r *http.Request)  {
	// load schema
	schema := loadSchema(&resolver.QueryResolver{}, &resolver.MutationResolver{})

	// graphql params
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		panic(err)
	}

	response := schema.Exec(context.Background(), params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Write(responseJSON)
}