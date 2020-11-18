package repository

import (
	"activities/domain/activity/entity"
	"context"
)

type Repository interface {
	GetAllActivities(ctx context.Context) (eActs []*entity.Activities, err error)
	GetActivityByID(ctx context.Context, req *entity.Activities) (eActs *entity.Activities, err error)
	InsertActivity(ctx context.Context, eAct *entity.Activities) (int64, error)
	UpdateActivity(ctx context.Context, eAct *entity.Activities) error
	DeleteActivity(ctx context.Context, eAct *entity.Activities) error
}
