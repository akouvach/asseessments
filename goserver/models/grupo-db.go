package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *DBModel) GrupoGetAll() ([]*Grupo, error) {

	var result bson.M
	var grupos []*Grupo
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	fmt.Println("llego a GetAll")
	// query := `select grupoId, nombre from grupos order by nombre`

	// rows, err := m.DB.QueryContext(ctx, query)
	rows := m.DB.Database("assessments").Collection("usuarios")
	err := rows.FindOne(context.TODO(), bson.D{{"title", "The Room"}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return grupos, err
		}
		panic(err)
	}
	// defer rows.Close()
	fmt.Println("ya hizo el select")

	// for rows.Next() {
	// 	var grupo Grupo
	// 	err := rows.Scan(
	// 		&grupo.GrupoId,
	// 		&grupo.Nombre,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	fmt.Println(grupo)
	// 	grupos = append(grupos, &grupo)

	// }

	return grupos, nil
}
