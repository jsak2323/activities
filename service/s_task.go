package service

// ALL BUSINESS LOGIC HERE

import (
	"activities/domain/activity/entity"
	"context"
	"errors"
)

type Activity struct {
	ID     int64
	Name   string
	Act    string
	Status int
}

func (s *svc) SelectAllActivities(ctx context.Context) ([]*Activity, error) {

	acts, err := s.act.GetAllActivities(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]*Activity, 0, len(acts))

	for _, v := range acts {
		resp = append(resp, &Activity{
			ID:     v.ID,
			Name:   v.Name,
			Act:    v.Act,
			Status: v.Status,
		})
	}

	return resp, err
}

func (s *svc) GetActivityByID(ctx context.Context, req *Activity) (*Activity, error) {
	eActs := new(entity.Activities)
	eActs.ID = req.ID
	eActs.Act = req.Act
	eActs.Status = req.Status

	acts, err := s.act.GetActivityByID(ctx, eActs)
	if err != nil {
		return nil, err
	}
	if acts == nil {
		return nil, errors.New("query failed")
	}

	resp := &Activity{
		ID:     acts.ID,
		Name:   acts.Name,
		Act:    acts.Act,
		Status: acts.Status,
	}

	return resp, err

}

func (s *svc) InsertActivity(ctx context.Context, req *Activity) (act *Activity, err error) {

	eActs := new(entity.Activities)
	eActs.Name = req.Name
	eActs.Act = req.Act
	eActs.Status = req.Status

	acts, err := s.act.InsertActivity(ctx, eActs)
	if err != nil {
		return nil, err
	}

	resp := &Activity{
		ID: acts,
	}

	return resp, err
}

func (s *svc) UpdateActivity(ctx context.Context, req *Activity) error {
	eActs := new(entity.Activities)
	eActs.ID = req.ID
	eActs.Name = req.Name
	eActs.Act = req.Act
	eActs.Status = req.Status

	err := s.act.UpdateActivity(ctx, eActs)
	if err != nil {
		return err
	}

	return err
}

func (s *svc) DeleteActivity(ctx context.Context, req *Activity) error {
	eActs := new(entity.Activities)
	eActs.ID = req.ID

	err := s.act.DeleteActivity(ctx, eActs)
	if err != nil {
		return err
	}

	return err
}
