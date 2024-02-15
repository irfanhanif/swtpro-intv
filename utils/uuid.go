package utils

import "github.com/google/uuid"

//go:generate mockgen -source=uuid.go -destination=mock/uuid.go -package=mock

type IUUID interface {
	New() uuid.UUID
}

type UUID struct{}

func NewUUID() *UUID {
	return &UUID{}
}

func (u *UUID) New() uuid.UUID {
	return u.New()
}
