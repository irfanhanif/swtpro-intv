package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/irfanhanif/swtpro-intv/entity"
	"github.com/irfanhanif/swtpro-intv/repository"
	"github.com/irfanhanif/swtpro-intv/valueobj"
)

type userUpdater struct {
	repo repository.IUpdateUserByID
}

func NewUserUpdater(repo repository.IUpdateUserByID) *userUpdater {
	return &userUpdater{repo: repo}
}

func (u *userUpdater) UpdateUserByID(ctx context.Context, id uuid.UUID, updateData valueobj.UserUpdateData) error {
	var errs []error

	if updateData.FullName != nil {
		fullNamePtr := updateData.FullName
		if fullNameErrs := entity.CheckFullName(*fullNamePtr); fullNameErrs != nil {
			errs = append(errs, fullNameErrs...)
		}
	}

	if updateData.PhoneNumber != nil {
		phoneNumberPtr := updateData.PhoneNumber
		if phoneNumberErrs := entity.CheckPhoneNumber(*phoneNumberPtr); phoneNumberErrs != nil {
			errs = append(errs, phoneNumberErrs...)
		}
	}

	errorsToStrings := func(errs []error) []string {
		result := []string{}
		for _, err := range errs {
			result = append(result, err.Error())
		}
		return result
	}

	if len(errs) > 0 {
		errFields := &ErrFields{Errs: []string{}}
		errFields.Errs = errorsToStrings(errs)
		return errFields
	}

	err := u.repo.UpdateUserByID(ctx, id, updateData)
	if errors.Is(err, repository.ErrNoRows) {
		return ErrNotFound
	}
	if errors.Is(err, repository.ErrPhoneNumberConflict) {
		return ErrPhoneNumberConflict
	}
	if err != nil {
		return err
	}

	return nil
}
