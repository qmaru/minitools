//go:build go1.27

package uuid

import (
	"uuid"
)

type UUIDBasic struct{}

type Version uint

const (
	Version4 Version = 4
	Version7 Version = 7
)

func New() *UUIDBasic {
	return new(UUIDBasic)
}

func (u *UUIDBasic) Generate(version Version) (*uuid.UUID, error) {
	var uid uuid.UUID

	switch version {
	case Version4:
		uid = uuid.NewV4()
	case Version7:
		uid = uuid.NewV7()
	default:
		uid = uuid.NewV4()
	}

	return &uid, nil
}
