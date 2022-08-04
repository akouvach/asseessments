package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *DBModel) PaisGetAll(db string) ([]*Pais, error) {

	var paises []*Pais
	fmt.Println("llego a GetAll")
	// query := `select paisId, nombre from paises order by paisId`

	// rows, err := m.DB.QueryContext(ctx, query)
	coll := m.DB.Database(db).Collection("geo_paises")
	title := "Kouvach"

	// var result bson.M
	// err = coll.Find(context.TODO(), bson.D{{"apellido", title}}).Decode(&result)
	rows, err := coll.Find(context.TODO(), bson.D{})
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return paises, err
	}
	if err != nil {
		panic(err)
	}
	defer rows.Close(context.TODO())

	if err = rows.All(context.TODO(), &paises); err != nil {
		return paises, err
	}

	for _, pais := range paises {
		fmt.Println(pais)
	}
	// jsonData, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%s\n", jsonData)

	// rows, err := coll.Find()
	// if err != nil {
	// 	return nil, err
	// }
	// defer rows.Close()
	// fmt.Println("ya hizo el select")
	// var paises []*Pais

	// for rows.Next() {
	// 	var pais Pais
	// 	err := rows.Scan(
	// 		&pais.PaisId,
	// 		&pais.Nombre,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	fmt.Println(pais)
	// 	paises = append(paises, &pais)

	// }

	return paises, nil
}
