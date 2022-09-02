package handler

import "github.com/graphql-go/graphql"

var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        ShelfType,
			Description: "Create new shelf",
			Args: graphql.FieldConfigArgument{
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

				}

				var shelf Shelf
				shelf.Room = params.Args["room"].(string)
				shelf.Location = params.Args["location"].(string)

				var newID int64
				sqlStatement := `INSERT INTO shelf(Room, Location) VALUES ($1, $2) RETURNING ID`
				err = db.QueryRow(sqlStatement, &shelf.Room, &shelf.Location).Scan(&newID)
				if err != nil {

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

				}

				var shelf Shelf
				sqlStatement := `UPDATE shelf SET room=$1, location=$2 WHERE id=$3 RETURNING id, room, location`
				err = db.QueryRow(sqlStatement, params.Args["room"].(string), params.Args["location"].(string), params.Args["id"].(int64)).Scan(&shelf.ID, &shelf.Room, &shelf.Location)
				if err != nil {

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

				}

				var shelf Shelf
				sqlStatement := `DELETE FROM shelf WHERE id=$1 RETURNING id`
				err = db.QueryRow(sqlStatement, params.Args["id"].(int64)).Scan(&shelf.ID)
				if err != nil {

				}

				return shelf, nil
			},
		},
	},
})
