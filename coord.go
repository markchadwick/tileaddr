package tileaddr

import (
	"errors"
	"io"
)

var ErrUnsupportedResolution = errors.New("Unsupported Resolution")

type Coord [2]int64

func (c Coord) WriteTo(w io.ByteWriter, zoom Zoom) error {
	return ErrUnsupportedResolution
}
