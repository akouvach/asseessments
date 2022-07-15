package models

import (
	"context"
	"fmt"
	"time"
)

func (m *DBModel) GrupoGetAll() ([]*Grupo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("llego a GetAll")
	query := `select grupoId, nombre from grupos order by nombre`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println("ya hizo el select")
	var grupos []*Grupo

	for rows.Next() {
		var grupo Grupo
		err := rows.Scan(
			&grupo.GrupoId,
			&grupo.Nombre,
		)
		if err != nil {
			return nil, err
		}
		fmt.Println(grupo)
		grupos = append(grupos, &grupo)

	}

	return grupos, nil
}
