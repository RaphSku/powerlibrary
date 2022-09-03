package handler

import (
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"shelf": &graphql.Field{
				Type:        ShelfType,
				Description: "Get shelf by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					db, err := ConnectToPSQL()
					if err != nil {
						return nil, err
					}

					id, ok := p.Args["id"].(int)
					if !ok {
						return nil, ArgumentMissing{"id"}
					}

					sqlStatement := `SELECT * FROM shelfs WHERE id=$1`
					row, err := db.Query(sqlStatement, id)
					if err != nil {
						return nil, err
					}
					defer row.Close()

					var shelf Shelf
					for row.Next() {
						err = row.Scan(&shelf.ID, &shelf.Location, &shelf.Location)
						if err != nil {
							return nil, err
						}
					}

					err = row.Err()
					if err != nil {
						return nil, err
					}

					return shelf, nil
				},
			},
			"shelfs": &graphql.Field{
				Type:        graphql.NewList(ShelfType),
				Description: "Get a list of shelfs",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					db, err := ConnectToPSQL()
					if err != nil {
						return nil, err
					}

					sqlStatement := `SELECT * FROM shelfs`
					rows, err := db.Query(sqlStatement)
					if err != nil {
						return nil, err
					}
					defer rows.Close()

					var shelfs Shelfs
					for rows.Next() {
						var shelf Shelf

						err = rows.Scan(&shelf.ID, &shelf.Location, &shelf.Room)
						if err != nil {
							return nil, err
						}

						shelfs = append(shelfs, &shelf)
					}

					err = rows.Err()
					if err != nil {
						return nil, err
					}

					return shelfs, nil
				},
			},
		},
	},
)
