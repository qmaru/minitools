package uuid

import (
	"errors"

	gouuid "github.com/gofrs/uuid/v5"
)

type UUIDSuiteBasic struct{}

type Option struct {
	Namespace gouuid.UUID
	Name      string
}

type Version int

const (
	Version1 Version = iota + 1
	Version4
	Version5
	Version7
)

func New() *UUIDSuiteBasic {
	return new(UUIDSuiteBasic)
}

func (u *UUIDSuiteBasic) SetNamespace(keywrod []byte) (gouuid.UUID, error) {
	return gouuid.FromBytes(keywrod)
}

func (u *UUIDSuiteBasic) Generate(version Version, option *Option) (string, error) {
	var uid gouuid.UUID
	var err error

	switch version {
	case Version1:
		uid, err = gouuid.NewV1()
	case Version4:
		uid, err = gouuid.NewV4()
	case Version5:
		if option == nil {
			return "", errors.New("option required for v5")
		}
		if option.Namespace == gouuid.Nil {
			return "", errors.New("invalid namespace for v5")
		}
		uid = gouuid.NewV5(option.Namespace, option.Name)
	case Version7:
		uid, err = gouuid.NewV7()
	default:
		uid, err = gouuid.NewV4()
	}

	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
