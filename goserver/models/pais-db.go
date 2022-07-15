package models

import (
	"context"
	"fmt"
	"time"
)

func (m *DBModel) PaisGetAll() ([]*Pais, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("llego a GetAll")
	query := `select paisId, nombre from paises order by paisId`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println("ya hizo el select")
	var paises []*Pais

	for rows.Next() {
		var pais Pais
		err := rows.Scan(
			&pais.PaisId,
			&pais.Nombre,
		)
		if err != nil {
			return nil, err
		}
		fmt.Println(pais)
		paises = append(paises, &pais)

	}

	return paises, nil
}
