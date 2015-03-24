package tileaddr

import (
	"github.com/google/gofuzz"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tile Addresses", func() {

	fuzz := fuzz.New()

	Context("at zoom zero", func() {

		It("should address zero for any values", func() {
			var (
				lng float64
				lat float64
			)
			for i := 0; i < 100; i++ {
				fuzz.Fuzz(&lng)
				fuzz.Fuzz(&lat)
				addr := Get(0, lng, lat)
				Expect(addr).To(Equal(uint64(0)))
			}
		})

	})

})
