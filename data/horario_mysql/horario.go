package horario_mysql

import (
	repo "agendamiento/data"
	model "agendamiento/model"
	"context"
	"database/sql"
	"fmt"
)

func NewSQLHorario(Conn *sql.DB) repo.Horario {
	return &mysqlHorario{
		Conn: Conn,
	}
}

type mysqlHorario struct {
	Conn *sql.DB
}

func (m *mysqlHorario) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Horario, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*model.Horario, 0)
	for rows.Next() {
		data := new(model.Horario)

		err := rows.Scan(
			&data.IDtutoria,
			&data.IDtutor,
			&data.NombreMateria,
			&data.Fecha,
			&data.HoraInicio,
			&data.HoraFinal,
			&data.Cupos,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlHorario) Fetch(ctx context.Context, num int64) ([]*model.Horario, error) {
	query := "Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlHorario) GetByID(ctx context.Context, IDtutoria int64) (*model.Horario, error) {
	query := "Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario where IDtutoria=?"
	rows, err := m.fetch(ctx, query, IDtutoria)
	if err != nil {
		return nil, err
	}

	fmt.Print("Select IDtutoria, IDtutor, NombreMateria, Fecha, HoraInicio, HoraFinal, Cupos From Horario where IDtutoria=")
	fmt.Println(IDtutoria)

	payload := &model.Horario{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, model.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlHorario) Create(ctx context.Context, p *model.Horario) (int64, error) {
	query := "Insert Horario SET IDtutoria=?, IDtutor=?, NombreMateria=?, Fecha=?, HoraInicio=?, HoraFinal=?, Cupos=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.IDtutoria, p.IDtutor, p.NombreMateria, p.Fecha, p.HoraInicio, p.HoraFinal, p.Cupos)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlHorario) Update(ctx context.Context, p *model.Horario) (*model.Horario, error) {
	query := "Update Horario set IDtutor=?, NombreMateria=?, Fecha=?, HoraInicio=?, HoraFinal=?, Cupos=? where IDtutoria=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.IDtutor,
		p.NombreMateria,
		p.Fecha,
		p.HoraInicio,
		p.HoraFinal,
		p.Cupos,
		p.IDtutoria,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlHorario) Delete(ctx context.Context, IDtutoria int64) (bool, error) {
	query := "Delete From Horario Where IDtutoria=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, IDtutoria)
	if err != nil {
		return false, err
	}
	return true, nil
}
