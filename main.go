package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// The query type defines the Top-Level API types.
var queryTopLevelTypes = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Wurld", nil
			},
		},
	},
})

// The schema defines the query schema and the mutation
// (optional, not included here) schema.
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryTopLevelTypes,
})

func main() {

	http.Handle("/graphql", handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	}))

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	fmt.Println("Go Play at http://localhost:3030")
	http.ListenAndServe(":3030", nil)
}
