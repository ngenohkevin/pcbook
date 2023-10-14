package service

import (
	"github.com/ngenohkevin/pcbook/pb"
	"sync"
)

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	//	save the laptop to store
	Save(laptop *pb.Laptop) error
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewInMemoryLaptopStore returns a new InMemoryLaptopStore
func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{data: make(map[string]*pb.Laptop)}
}
