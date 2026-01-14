package uuid

import (
	"errors"

	gouuid "github.com/gofrs/uuid/v5"
)

type UUIDBasic struct{}

type Option struct {
	Namespace []byte
	Name      string
}

type Version uint

const (
	Version1 Version = iota + 1
	Version4
	Version5
	Version7
)

func New() *UUIDBasic {
	return new(UUIDBasic)
}

func (u *UUIDBasic) setNamespace(keywrod []byte) (gouuid.UUID, error) {
	if len(keywrod) != 16 {
		return gouuid.Nil, errors.New("invalid namespace length")
	}
	return gouuid.FromBytes(keywrod)
}

func (u *UUIDBasic) Generate(version Version, option *Option) (string, error) {
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
		if len(option.Namespace) != 16 {
			return "", errors.New("invalid namespace for v5")
		}

		name := option.Name
		namespace, err := u.setNamespace(option.Namespace)
		if err != nil {
			return "", err
		}
		uid = gouuid.NewV5(namespace, name)
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
