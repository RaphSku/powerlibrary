package handler

import "github.com/graphql-go/graphql"

var ShelfType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Shelf",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"room": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
