package storage

import (
	"fmt"
	"gitTest/bins"
)

type Storage struct {
	bins map[string]*bins.Bin
}

func NewStorage() *Storage {
	return &Storage{
		bins: make(map[string]*bins.Bin),
	}
}

func (s *Storage) CreateBin(bin *bins.Bin) error {
	if bin == nil {
		return fmt.Errorf("Bin is nil")
	}
	s.bins[bin.Name] = bin
	return nil
}

func (s *Storage) LookUpBun(name string) (*bins.Bin, error) {
	bin, ok := s.bins[name]

	if !ok {
		return nil, fmt.Errorf("Bin not found %s", name)
	}
	return bin, nil
}

func (s *Storage) DeleteBin(name string) error {
	if _, ok := s.bins[name]; !ok {
		return fmt.Errorf("Bin not found %s", name)
	}
	delete(s.bins, name)
	return nil
}

func (s *Storage) ListBinNames() []string {
	var names []string
	for name := range s.bins {
		names = append(names, name)
	}
	return names
}
