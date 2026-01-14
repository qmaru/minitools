package uuid

import (
	"errors"

	gouuid "github.com/gofrs/uuid/v5"
)

type UUIDBasic struct{}

type Version uint

type Option func(*opts) error

type opts struct {
	Namespace *gouuid.UUID
	Name      string
}

const (
	Version1 Version = 1
	Version4 Version = 4
	Version5 Version = 5
	Version7 Version = 7
)

func New() *UUIDBasic {
	return new(UUIDBasic)
}

func WithNamespace(ns gouuid.UUID) Option {
	return func(o *opts) error {
		o.Namespace = &ns
		return nil
	}
}

func WithNamespaceBytes(b []byte) Option {
	return func(o *opts) error {
		if len(b) != 16 {
			return errors.New("invalid namespace length")
		}
		ns, err := gouuid.FromBytes(b)
		if err != nil {
			return err
		}
		o.Namespace = &ns
		return nil
	}
}

func WithName(name string) Option {
	return func(o *opts) error {
		o.Name = name
		return nil
	}
}

func (u *UUIDBasic) Generate(version Version, options ...Option) (string, error) {
	var uid gouuid.UUID
	var err error

	cfg := &opts{}
	for _, opt := range options {
		if opt == nil {
			continue
		}
		if err := opt(cfg); err != nil {
			return "", err
		}
	}

	switch version {
	case Version1:
		uid, err = gouuid.NewV1()
	case Version4:
		uid, err = gouuid.NewV4()
	case Version5:
		if cfg.Namespace == nil {
			return "", errors.New("namespace required for v5")
		}
		uid = gouuid.NewV5(*cfg.Namespace, cfg.Name)
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
