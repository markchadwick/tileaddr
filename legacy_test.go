package tileaddr

import (
	"log"

	"github.com/markchadwick/morton"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Legacy Equality", func() {

	var depth uint8 = 8
	var x float64 = 232.2235
	var y float64 = 12.0

	It("should get the same results", func() {
		legacy := encodeAddress(depth, [2]float64{x, y})
		current := morton.Enc32(uint32(x), uint32(y))

		log.Printf("------------------------------------")
		log.Printf("Legacy:  %08x", legacy)
		log.Printf("Current: %08x", current)
	})
})
