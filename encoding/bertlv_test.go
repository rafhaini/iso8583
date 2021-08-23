package encoding

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBerTLVTag_Encode(t *testing.T) {
        res, err := EBCDIC.Encode([]byte{0x5F, 0x2A})
        require.NoError(t, err)
        require.Equal(t, []byte("5F2A"), res)
}

// func TestBerTLVTag(t *testing.T) {
// 	t.Run("Decode", func(t *testing.T) {
// 		res, read, err := EBCDIC.Decode([]byte{0x12, 0x34}, 2)
// 
// 		require.NoError(t, err)
// 		require.Equal(t, []byte{0x12, 0x94}, res)
// 		require.Equal(t, 2, read)
// 
// 	})
// 
// 	t.Run("Encode", func(t *testing.T) {
// 		res, err := EBCDIC.Encode([]byte{0x12, 0x94})
// 
// 		require.NoError(t, err)
// 		require.Equal(t, []byte{0x12, 0x34}, res)
// 
// 	})
// }
