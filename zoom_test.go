package tileaddr

import (
	"bytes"
	"io"

	"github.com/google/gofuzz"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Zoom encoding", func() {

	fuzz := fuzz.New()
	var buf *bytes.Buffer

	BeforeEach(func() {
		buf = &bytes.Buffer{}
	})

	Context("when writing", func() {

		It("should write zero as 0x00", func() {
			Zoom(0).WriteTo(buf)
			bs := buf.Bytes()
			Expect(bs).To(HaveLen(1))
			Expect(bs[0]).To(Equal(uint8(0x00)))
		})

		It("should write 255 as 0xff", func() {
			Zoom(255).WriteTo(buf)
			bs := buf.Bytes()
			Expect(bs).To(HaveLen(1))
			Expect(bs[0]).To(Equal(uint8(0xff)))
		})

		It("should produce orderable values", func() {
			var za, zb uint8
			for i := 0; i < 100; i++ {
				fuzz.Fuzz(&za)
				fuzz.Fuzz(&zb)
				ba := zoomBytes(Zoom(za))
				bb := zoomBytes(Zoom(zb))

				if za > zb {
					Expect(bytes.Compare(ba, bb)).To(Equal(1))
				} else if za < zb {
					Expect(bytes.Compare(ba, bb)).To(Equal(-1))
				} else {
					Expect(bytes.Compare(ba, bb)).To(Equal(0))
				}
			}
		})

	})

	Context("when reading", func() {

		It("should read 0xff as 255", func() {
			z, _ := ReadZoom(bytes.NewBuffer([]byte{0xff}))
			Expect(z).To(Equal(Zoom(255)))
		})

		It("should read 0x00 as 0", func() {
			z, _ := ReadZoom(bytes.NewBuffer([]byte{0x00}))
			Expect(z).To(Equal(Zoom(0)))
		})

		It("should propigate any read error", func() {
			_, err := ReadZoom(&bytes.Buffer{})
			Expect(err).To(Equal(io.EOF))
		})

	})

	Context("resolution", func() {

		It("should be 1 for zoom 0", func() {
			Expect(Zoom(0).resolution()).To(Equal(1))
		})

		It("should be 2 for zoom 1", func() {
			Expect(Zoom(1).resolution()).To(Equal(2))
		})

		It("should be 4 for zoom 2", func() {
			Expect(Zoom(2).resolution()).To(Equal(4))
		})

		It("should be 16k for zoom 14", func() {
			Expect(Zoom(14).resolution()).To(Equal(16384))
		})

	})

})

func zoomBytes(z Zoom) []byte {
	buf := &bytes.Buffer{}
	err := z.WriteTo(buf)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
