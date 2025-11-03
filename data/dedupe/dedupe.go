package dedupe

import (
	"fmt"
	"hash/fnv"

	"golang.org/x/sync/singleflight"
)

type HashFunc func(parts ...any) string

type Option[T any] func(*SF[T])

type SF[T any] struct {
	group    *singleflight.Group
	hashFunc HashFunc
}

type TResult[T any] struct {
	Val    T
	Err    error
	Shared bool
}

func defaultHashKey(parts ...any) string {
	h := fnv.New64a()
	for _, part := range parts {
		_, _ = h.Write(fmt.Appendf(nil, "%v|", part))
	}
	return fmt.Sprintf("%x", h.Sum64())
}

func NewT[T any](opts ...Option[T]) *SF[T] {
	sf := &SF[T]{
		group:    &singleflight.Group{},
		hashFunc: defaultHashKey,
	}
	for _, opt := range opts {
		opt(sf)
	}
	return sf
}

func New(opts ...Option[any]) *SF[any] {
	return NewT(opts...)
}

func WithHash[T any](hashFunc HashFunc) Option[T] {
	return func(s *SF[T]) {
		if hashFunc != nil {
			s.hashFunc = hashFunc
		}
	}
}

func (s *SF[T]) HashKey(parts ...any) string {
	return s.hashFunc(parts...)
}

func (s *SF[T]) Forget(key string) {
	s.group.Forget(key)
}

func (s *SF[T]) Do(key string, fn func() (any, error)) (any, error) {
	v, err, _ := s.group.Do(key, fn)
	return v, err
}

func (s *SF[T]) DoChan(key string, fn func() (any, error)) <-chan singleflight.Result {
	return s.group.DoChan(key, fn)
}

func (s *SF[T]) DoT(key string, fn func() (T, error)) (T, error) {
	v, err, _ := s.group.Do(key, func() (any, error) {
		return fn()
	})
	if err != nil {
		var zero T
		return zero, err
	}
	if v == nil {
		var zero T
		return zero, nil
	}
	result, ok := v.(T)
	if !ok {
		var zero T
		return zero, fmt.Errorf("type assertion failed: expected %T, got %T", zero, v)
	}
	return result, nil
}

func (s *SF[T]) DoChanT(key string, fn func() (T, error)) <-chan TResult[T] {
	ch := make(chan TResult[T], 1)
	go func() {
		defer close(ch)
		v, err, shared := s.group.Do(key, func() (any, error) {
			return fn()
		})
		var result T
		if err == nil && v != nil {
			result = v.(T)
		}
		ch <- TResult[T]{Val: result, Err: err, Shared: shared}
	}()
	return ch
}
