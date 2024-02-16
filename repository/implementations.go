package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/lib/pq"
	"github.com/pkg/errors"
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

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.IUser, error) {
	var userModel UserModel

	query := `select id, phone_number, password, full_name from "user" where phone_number = $1`
	row := r.Db.QueryRow(query, phoneNumber)
	switch err := row.Scan(
		&userModel.ID,
		&userModel.PhoneNumber,
		&userModel.Password,
		&userModel.FullName,
	); errors.Cause(err) {
	case nil:
		return entity.NewUserFactory(nil).NewUserWithID(userModel.ID, userModel.PhoneNumber, userModel.Password, userModel.FullName), nil
	case sql.ErrNoRows:
		return nil, ErrNoRows
	default:
		return nil, err
	}
}

func (r *Repository) IncrementLoginCount(ctx context.Context, userID uuid.UUID) error {
	query := `update "user" set login_count = login_count + 1 where id = $1`
	_, err := r.Db.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (entity.IUser, error) {
	var userModel UserModel

	query := `select id, phone_number, password, full_name from "user" where id = $1`
	row := r.Db.QueryRow(query, id)
	switch err := row.Scan(
		&userModel.ID,
		&userModel.PhoneNumber,
		&userModel.Password,
		&userModel.FullName,
	); errors.Cause(err) {
	case nil:
		return entity.NewUserFactory(nil).NewUserWithID(userModel.ID, userModel.PhoneNumber, userModel.Password, userModel.FullName), nil
	case sql.ErrNoRows:
		return nil, ErrNoRows
	default:
		return nil, err
	}
}
