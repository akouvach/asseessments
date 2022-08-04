package models

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func (m *DBModel) UsuarioGetOne(usuarioId int64) (*Usuario, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	query := `select usuarioid, apellido,nombre,
// 	email, pass, paisid
// 	from usuarios
// 	where usuarioid = $1`

// 	row := m.DB.QueryRowContext(ctx, query, usuarioId)

// 	var usuario Usuario

// 	err := row.Scan(
// 		&usuario.UsuarioId,
// 		&usuario.Apellido,
// 		&usuario.Nombre,
// 		&usuario.Email,
// 		// &usuario.FechaRegistro,
// 		&usuario.Pass,
// 		&usuario.PaisId,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	// agrego los grupos que tiene este usuario

// 	query = `select usuarioId, grupoId
// 	from gruposintegrantes where usuarioId = $1 and activo`

// 	rows, _ := m.DB.QueryContext(ctx, query, usuarioId)
// 	defer rows.Close()

// 	var misGrupos []GrupoIntegrante

// 	for rows.Next() {
// 		var g GrupoIntegrante
// 		err := rows.Scan(
// 			&g.UsuarioId,
// 			&g.GrupoId,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fmt.Println(g)
// 		misGrupos = append(misGrupos, g)

// 	}

// 	// usuario.Grupos = misGrupos

// 	return &usuario, nil

// }

// func (m *DBModel) UsuarioGetAll() ([]*Usuario, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	fmt.Println("llego a GetAll")
// 	query := `select nombre, apellido, email, usuarioId
// 	from usuarios order by email`

// 	rows, err := m.DB.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	fmt.Println("ya hizo el select")
// 	var usuarios []*Usuario

// 	for rows.Next() {
// 		var usu Usuario
// 		err := rows.Scan(
// 			&usu.Nombre,
// 			&usu.Apellido,
// 			&usu.Email,
// 			&usu.UsuarioId,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		fmt.Println(usu)
// 		usuarios = append(usuarios, &usu)

// 	}

// 	return usuarios, nil
// }

// func (m *DBModel) UsuarioPut(usu Usuario) (Usuario, error) {

// 	var query string
// 	var err error
// 	var lastid int64
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	fmt.Println("llego a Put")
// 	if usu.UsuarioId == 0 {
// 		fmt.Println("es un insert")
// 		lastid = 0
// 		/* es un insert */
// 		query = `insert into usuarios (nombre, apellido, email, pass, paisid)
// 		values ($1, $2, $3, $4, $5) returning usuarioid`
// 		row := m.DB.QueryRowContext(ctx, query,
// 			usu.Nombre,
// 			usu.Apellido,
// 			usu.Email,
// 			usu.Pass,
// 			usu.PaisId,
// 		)
// 		err = row.Scan(
// 			&lastid,
// 		)
// 		fmt.Println("id:", lastid)
// 		if err != nil {
// 			fmt.Println("error en el insert ", err)
// 			return usu, err
// 		}
// 		usu.UsuarioId = lastid
// 	} else {
// 		lastid = usu.UsuarioId
// 		/* es un update */
// 		query = `update usuarios set nombre=$2, apellido=$3, email=$4, pass=$5, paisid=$6
// 		where usuarioId=$1`
// 		_, err := m.DB.ExecContext(ctx, query,
// 			usu.UsuarioId,
// 			usu.Nombre,
// 			usu.Apellido,
// 			usu.Email,
// 			usu.Pass,
// 			usu.PaisId,
// 		)
// 		if err != nil {
// 			fmt.Println("error en el update ", err)
// 			return usu, err
// 		}
// 	}

// 	// lastInsertId := 0
// 	// err = db.QueryRow("INSERT INTO brands (name) VALUES($1) RETURNING id", name).Scan(&lastInsertId)

// 	fmt.Println("Resultado del insert:", lastid)
// 	fmt.Println("Resultado del usuario:", usu)

// 	return usu, nil
// }

// func (m *DBModel) UsuarioDelete(usuarioId int64) error {

// 	var query string
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	fmt.Println("llego a Delete")
// 	if usuarioId > 0 {
// 		query = `delete from usuarios where usuarioid = $1`
// 		_, err := m.DB.ExecContext(ctx, query, usuarioId)
// 		if err != nil {
// 			fmt.Println("error en el delete", err)
// 			return err
// 		}
// 	}

// 	return nil
// }
