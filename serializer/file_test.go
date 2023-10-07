package serializer

import (
	"github.com/golang/protobuf/proto"
	"github.com/ngenohkevin/pcbook/pb"
	"github.com/ngenohkevin/pcbook/sample"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobuffFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}
