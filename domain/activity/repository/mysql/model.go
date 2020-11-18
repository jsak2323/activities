package mysql

import (
	"activities/domain/activity/entity"

	"gopkg.in/nullbio/null.v6"
)

type Activities struct {
	ID     null.Int64 `db:"id,omitempty"`
	Name   string     `db:"name"`
	Act    string     `db:"act"`
	Status int        `db:"status"`
}

func (act Activities) ToEntity() *entity.Activities {
	return &entity.Activities{
		ID:     act.ID.Int64,
		Name:   act.Name,
		Act:    act.Act,
		Status: act.Status,
	}
}

// ToEntities ...
func ToEntities(acts []Activities) []*entity.Activities {
	var eActs []*entity.Activities

	for _, v := range acts {
		eActs = append(eActs, v.ToEntity())
	}

	return eActs
}
