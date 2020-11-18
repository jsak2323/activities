package mysql

import (
	"activities/domain/activity/entity"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"gopkg.in/nullbio/null.v6"
)

// MySQL ...
type MySQL struct {
	db *sqlx.DB
}

// New ...
func New(db *sqlx.DB) *MySQL {
	return &MySQL{db}
}

func (m *MySQL) GetAllActivities(ctx context.Context) (eActs []*entity.Activities, err error) {

	var activities []Activities

	query := "SELECT id, name, act, status FROM activities"

	rows, err := m.db.Query(query)
	if err != nil {
		return eActs, err
	}

	err = sqlx.StructScan(rows, &activities)
	if err != nil {
		return eActs, err
	}
	rows.Close()

	eActs = ToEntities(activities)

	return eActs, err
}

func (m *MySQL) GetActivityByID(ctx context.Context, req *entity.Activities) (eActs *entity.Activities, err error) {

	var activities Activities

	query := "SELECT id, name, act, status FROM activities WHERE id = ?"

	err = m.db.GetContext(ctx, &activities, query, req.ID)
	// rows, err := m.db.Queryx(query, req.ID)
	if err != nil {
		return eActs, err
	}

	// defer rows.Close()

	// for rows.Next() {
	// 	rows.StructScan(&activities)

	// }
	eActs = activities.ToEntity()

	return eActs, err
}

func (m *MySQL) InsertActivity(ctx context.Context, eAct *entity.Activities) (int64, error) {

	// name, act string, status int
	mAct := new(Activities)
	mAct.Name = eAct.Name
	mAct.Act = eAct.Act
	mAct.Status = eAct.Status

	query := `INSERT INTO 
				activities (
					name,
					act, 
					status
				) VALUES (
					:name, 
					:act, 
					:status
				);`

	var result sql.Result

	result, err := m.db.NamedExecContext(ctx, query, mAct)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return ID, err
}

func (m *MySQL) UpdateActivity(ctx context.Context, eAct *entity.Activities) error {

	// name, act string, status int
	mAct := new(Activities)
	mAct.ID = null.Int64From(eAct.ID)
	mAct.Name = eAct.Name
	mAct.Act = eAct.Act
	mAct.Status = eAct.Status

	query := `UPDATE 
					activities 
				SET 
					name = :name, 
					act = :act, 
					status = :status 
				WHERE 
					id = :id`

	var result sql.Result
	result, err := m.db.NamedExecContext(ctx, query, mAct)
	if err != nil {
		return err
	}

	res, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if res == 0 {
		return errors.New("no rows affected")
	}

	return err
}

func (m *MySQL) DeleteActivity(ctx context.Context, eAct *entity.Activities) error {
	mAct := new(Activities)
	mAct.ID = null.Int64From(eAct.ID)

	query := "DELETE FROM activities WHERE id = :id"

	var result sql.Result
	result, err := m.db.NamedExecContext(ctx, query, mAct)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return err
}
