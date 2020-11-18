package service

import (
	activity "activities/domain/activity/repository"
	"context"
)

type svc struct {
	act activity.Repository
}

// New instantiate service object which include all business logic
func New(
	acts activity.Repository,
) Service {
	return &svc{
		act: acts,
	}
}

type Service interface {
	SelectAllActivities(ctx context.Context) ([]*Activity, error)
	GetActivityByID(ctx context.Context, req *Activity) (*Activity, error)
	InsertActivity(ctx context.Context, req *Activity) (*Activity, error)
	UpdateActivity(ctx context.Context, req *Activity) error
	DeleteActivity(ctx context.Context, req *Activity) error
}
