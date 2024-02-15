package repository

import (
	"context"
	"github.com/irfanhanif/swtpro-intv/entity"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateNewUser(ctx context.Context, ua entity.IUserAuthentication, up entity.IUserProfile) error {
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query1 := `insert into user_profile (id, full_name) values($1, $2)`
	_, err = tx.Exec(query1,
		up.ID(),
		up.FullName(),
	)
	if err != nil {
		return err
	}

	query2 := `insert into user_authentication (id, user_id, phone_number, password) values ($1, $2, $3, $4)`
	_, err = tx.Exec(query2,
		up.ID(),
		up.FullName(),
	)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
