package service

import (
	"github.com/ngenohkevin/pcbook/pb"
	"github.com/ngenohkevin/pcbook/sample"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  LaptopStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
		},
	}
}
