package tileaddr

import (
	"io"
)

type Zoom uint8

// Write the zoom as 8 bits to an underlying buffer
func (z Zoom) WriteTo(w io.ByteWriter) error {
	return w.WriteByte(uint8(z))
}

// Read the zoom from an underlying buffer, advancing the input 8 bits. Any
// error thrown while reading the buffer will be propagated to the calling
// function.
func ReadZoom(r io.ByteReader) (Zoom, error) {
	if b, err := r.ReadByte(); err != nil {
		return 0, err
	} else {
		return Zoom(b), nil
	}
}

// Resolution in a single dimension of this zoom. Size of the range of values
// from the visible minimum to maximum.
func (z Zoom) resolution() int {
	return 1 << z
}
