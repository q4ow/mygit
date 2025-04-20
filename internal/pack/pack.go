package pack

import (
	"encoding/binary"
	"io"

	"github.com/q4ow/mygit/internal/objects"
)

func Pack(objects []*objects.Object, w io.Writer) error {
	for _, obj := range objects {
		if _, err := w.Write([]byte{0x00, 0x00, 0x00, 0x00}); err != nil {
			return err
		}

		size := make([]byte, 4)
		binary.BigEndian.PutUint32(size, uint32(len(obj.Content)))

		if _, err := w.Write([]byte{0x74}); err != nil {
			return err
		}

		if _, err := w.Write(size); err != nil {
			return err
		}

		if _, err := w.Write(obj.Content); err != nil {
			return err
		}
	}
	return nil
}
