package storage

import (
	"go.uber.org/zap"
	"strconv"
)

type Storage struct {
	inner  map[string]string
	logger *zap.Logger
}

func NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("created new storage")

	return Storage{
		inner:  make(map[string]string),
		logger: logger,
	}, nil
}

func (r Storage) Set(key, value string) {
	r.inner[key] = value
	r.logger.Info("key set", zap.String("key", key), zap.String("value", r.inner[key]))
}

func (r Storage) Get(key string) string {
	r.logger.Info("key obtained")
	return r.inner[key]
}

func (r Storage) GetKind(key string) string {
	_, err := strconv.Atoi(key)
	s := "D"
	if err != nil {
		s = "S"
		return s
	}
	return s
}