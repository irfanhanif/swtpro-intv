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

func (r *Repository) CreateNewUser(ctx context.Context, user entity.IUser) error {
	return nil
}
