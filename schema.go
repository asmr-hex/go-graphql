package main

import (
	"github.com/graphql-go/graphql"
)

// Define Human Type
var HumanType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Human",
	Description: "Human character within 'Snow Crash'",
	Fields: graphql.Fields{
		"_id": &graphql.Field{
			// The NonNull is a type marker which specifies that
			// the Type of this field MUST be a non empty string.
			Type: graphql.NewNonNull(graphql.String),
			// The Description is a decription of this field.
			Description: "The id of a human",
			// The Resolve value must be a function with the following
			// signature. This function is responsible for actually
			// fetching the data for this field. The Resolve functions
			// are really the glue between the GraphQL engine and our
			// backend data storage solution.
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// typically, we want to ...
			},
		},
	},
})

// Define Machine Type
var MachineType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Machine",
	Description: "Machine character within 'Snow Crash'",
	Fields: graphql.Fields{
		"_id": &graphql.Field{},
	},
})

// Define Organization Type
var OrganizationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Organization",
	Description: "Organization within 'Snow Crash'",
	Fields: graphql.Fields{
		"_id": &graphql.Field{},
	},
})

// Define ExistentialPlane Type
var ExistentialPlaneType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "ExistentialPlane",
	Description: "An existential plane within 'Snow Crash'",
	Fields: graphql.Fields{
		"_id": &graphql.Field{},
	},
})

// Define GeoLocation Type
var GeoLocationType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoLocation",
	Description: "A Geo-Location within 'Snow Crash'",
	Fields: graphql.Fields{
		"_id": &graphql.Field{},
	},
})
