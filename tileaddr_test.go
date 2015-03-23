package tileaddr

import (
	"math"
	"strconv"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tile Address")
}

// Legacy iterative function which imperatively descends the address of a
// quadtree for a given address and depth
func encodeAddress(depth uint8, coord [2]float64) []byte {
	nBytes := int(math.Ceil(float64(depth) / 4))
	addr := make([]byte, nBytes)

	var (
		x = coord[0]
		y = coord[1]

		xMin float64 = 0
		xMax float64 = 256
		xMid float64 = 128

		yMin float64 = 0
		yMax float64 = 256
		yMid float64 = 128
	)

	iDepth := int(depth)
	for i := 0; i < iDepth; i++ {
		bitIdx := i * 2
		byteIdx := int(math.Floor(float64(bitIdx) / 8))
		bitOff := uint(bitIdx - (byteIdx * 8))

		// Write X bit
		if x >= xMid {
			addr[byteIdx] |= 0x01 << (7 - bitOff)
			xMin = xMid
		} else {
			xMax = xMid
		}
		xMid = (xMin / 2) + (xMax / 2)

		// Write Y bit
		if y >= yMid {
			addr[byteIdx] |= 0x01 << (6 - bitOff)
			yMin = yMid
		} else {
			yMax = yMid
		}
		yMid = (yMin / 2) + (yMax / 2)
	}

	return addr
}

func binStr(i uint64, bits int) string {
	s := strconv.FormatUint(i, 2)
	return strings.Repeat("0", bits-len(s)) + s
}
