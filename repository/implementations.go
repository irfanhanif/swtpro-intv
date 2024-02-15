package repository

import (
	"context"
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/lib/pq"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateNewUser(ctx context.Context, user entity.IUser) error {
	query := `insert into "user" (id, phone_number, password, full_name) values ($1, $2, $3, $4)`
	_, err := r.Db.Exec(query,
		user.ID(),
		user.PhoneNumber(),
		user.HashedPassword(),
		user.FullName(),
	)
	pqErr, ok := err.(*pq.Error)
	if ok && pqErr.Code == "23505" && pqErr.Constraint == "user_phone_number_unique" {
		return ErrPhoneNumberConflict
	}
	if err != nil {
		return err
	}

	return nil
}
