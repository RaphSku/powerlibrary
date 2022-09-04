package handler

import (
	"github.com/graphql-go/graphql"
)

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        ShelfType,
			Description: "Create new shelf",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"room": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"location": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				db, err := ConnectToPSQL()
				if err != nil {
					return nil, err
				}

				var shelf Shelf
				shelf.Name = params.Args["name"].(string)
				shelf.Room = params.Args["room"].(string)
				shelf.Location = params.Args["location"].(string)

				var newID int64
				sqlStatement := `INSERT INTO shelfs(Name, Room, Location) VALUES ($1, $2, $3) RETURNING ID`
				err = db.QueryRow(sqlStatement, &shelf.Name, &shelf.Room, &shelf.Location).Scan(&newID)
				if err != nil {
					return nil, err
				}

				shelf.ID = newID

				return shelf, nil
			},
		},
		"update": &graphql.Field{
			Type:        ShelfType,
			Description: "Update shelf",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"room": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"location": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				db, err := ConnectToPSQL()
				if err != nil {
					return nil, err
				}

				var shelf Shelf
				sqlStatement := `UPDATE shelfs SET name=$1, room=$2, location=$3 WHERE id=$4 RETURNING id, name, room, location`
				err = db.QueryRow(sqlStatement, params.Args["name"].(string), params.Args["room"].(string), params.Args["location"].(string),
					params.Args["id"].(int64)).Scan(&shelf.ID, &shelf.Name, &shelf.Room, &shelf.Location)
				if err != nil {
					return nil, err
				}

				return shelf, nil
			},
		},
		"delete": &graphql.Field{
			Type:        ShelfType,
			Description: "Delete shelf",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"room": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"location": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				db, err := ConnectToPSQL()
				if err != nil {
					return nil, err
				}

				var shelf Shelf
				sqlStatement := `DELETE FROM shelfs WHERE id=$1 RETURNING id`
				err = db.QueryRow(sqlStatement, params.Args["id"].(int64)).Scan(&shelf.ID)
				if err != nil {
					return nil, err
				}

				return shelf, nil
			},
		},
	},
})
