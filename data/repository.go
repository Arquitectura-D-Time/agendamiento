package data

import (
	"agendamiento/model"
	"context"
)

type Horario interface {
	Fetch(ctx context.Context, num int64) ([]*model.Horario, error)
	GetByID(ctx context.Context, id int64) (*model.Horario, error)
	Create(ctx context.Context, p *model.Horario) (int64, error)
	Update(ctx context.Context, p *model.Horario) (*model.Horario, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

type Agendadas interface {
	Fetch(ctx context.Context, num int64) ([]*model.Agendadas, error)
	GetByID(ctx context.Context, id int64) (*model.Agendadas, error)
	Create(ctx context.Context, p *model.Agendadas) (int64, error)
	Update(ctx context.Context, p *model.Agendadas) (*model.Agendadas, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
