package service

import "github.com/ngenohkevin/pcbook/pb"

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	//	save the laptop to store
	Save(laptop *pb.Laptop) error
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	data map[string]*pb.Laptop
}
