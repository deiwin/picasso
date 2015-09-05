package picasso_test

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	. "github.com/deiwin/picasso"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestImage string

const (
	GirlBeforeAMirror TestImage = "./test_images/picasso-girl_before_a_mirror.jpg"
	OldGuitarist      TestImage = "./test_images/picasso-the_old_guitarist.jpg"
	WomenOfAlgiers    TestImage = "./test_images/picasso-the_women_of_algiers.jpg"
	Bullfight         TestImage = "./test_images/picasso-bullfight.jpg"
	Composed          TestImage = "./test_images/composed.png"
)

func (i TestImage) read() image.Image {
	file, err := os.Open(string(i))
	Expect(err).NotTo(HaveOccurred())
	defer file.Close()

	image, _, err := image.Decode(file)
	Expect(err).NotTo(HaveOccurred())
	return image
}

var _ = Describe("Picasso", func() {
	Describe("Node", func() {
		It("draws the composed image", func() {
			i := HorizontalSplit{
				Ratio: 2,
				Top:   Picture{Bullfight.read()},
				Bottom: VerticalSplit{
					Ratio: 0.5,
					Left:  Picture{GirlBeforeAMirror.read()},
					Right: VerticalSplit{
						Ratio: 1,
						Left:  Picture{OldGuitarist.read()},
						Right: Picture{WomenOfAlgiers.read()},
					},
				},
			}.Draw(400, 600)

			composed := Composed.read()
			Expect(i.Bounds().Dx()).To(Equal(composed.Bounds().Dx()))
			Expect(i.Bounds().Dy()).To(Equal(composed.Bounds().Dy()))
			for x := 0; x < i.Bounds().Dx(); x++ {
				for y := 0; y < i.Bounds().Dy(); y++ {
					Expect(i.At(x, y)).To(Equal(composed.At(x, y)))
				}
			}
		})
	})
})
