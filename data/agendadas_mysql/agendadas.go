package agendadas_mysql

import (
	repo "agendamiento/data"
	model "agendamiento/model"
	"context"
	"database/sql"
)

func NewSQLHorario(Conn *sql.DB) repo.Agendadas {
	return &mysqlAgendadas{
		Conn: Conn,
	}
}

type mysqlAgendadas struct {
	Conn *sql.DB
}

func (m *mysqlAgendadas) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.Agendadas, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*model.Agendadas, 0)
	for rows.Next() {
		data := new(model.Agendadas)

		err := rows.Scan(
			&data.IDtutoria,
			&data.IDalumno,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlAgendadas) Fetch(ctx context.Context, num int64) ([]*model.Agendadas, error) {
	query := "Select IDtutoria, IDalumno From Agendadas limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlAgendadas) GetByID(ctx context.Context, IDtutoria int64) (*model.Agendadas, error) {
	query := "Select IDtutoria, IDalumno From Agendadas where IDtutoria=?"

	rows, err := m.fetch(ctx, query, IDtutoria)
	if err != nil {
		return nil, err
	}

	payload := &model.Agendadas{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, model.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlAgendadas) Create(ctx context.Context, p *model.Agendadas) (int64, error) {
	query := "Insert Agendadas SET IDtutoria=?, IDalumno=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.IDtutoria, p.IDalumno)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlAgendadas) Update(ctx context.Context, p *model.Agendadas) (*model.Agendadas, error) {
	query := "Update Agendadas set IDalumno=? where IDtutoria=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.IDalumno,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlAgendadas) Delete(ctx context.Context, IDtutoria int64) (bool, error) {
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
