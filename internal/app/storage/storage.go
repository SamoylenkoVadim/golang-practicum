package storage

import "errors"

type Storage struct {
	storage map[string]string
}

func New() *Storage {
	return &Storage{make(map[string]string)}
}

func (s *Storage) Save(key string, link string) error {
	if _, exist := s.storage[key]; exist {
		return errors.New("already exist")
	}
	s.storage[key] = link
	return nil
}

func (s *Storage) GetValue(key string) (string, bool) {
	if value, ok := s.storage[key]; ok {
		return value, true
	}
	return "", false
}
