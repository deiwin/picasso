package picasso_test

import (
	"image"
	_ "image/jpeg"
	"image/png"
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
	TopHeavy1         TestImage = "./test_images/top_heavy-1.png"
	TopHeavy2         TestImage = "./test_images/top_heavy-2.png"
	TopHeavy3         TestImage = "./test_images/top_heavy-3.png"
	TopHeavy4         TestImage = "./test_images/top_heavy-4.png"
)

func (i TestImage) read() image.Image {
	file, err := os.Open(string(i))
	Expect(err).NotTo(HaveOccurred())
	defer file.Close()

	image, _, err := image.Decode(file)
	Expect(err).NotTo(HaveOccurred())
	return image
}

func (i TestImage) write(data image.Image) {
	outfile, err := os.Create(string(i))
	Expect(err).NotTo(HaveOccurred())
	defer outfile.Close()

	png.Encode(outfile, data)
}

var _ = Describe("Picasso", func() {
	ExpectToEqualTestImage := func(i image.Image, testImage TestImage) {
		composed := testImage.read()
		Expect(i.Bounds().Dx()).To(Equal(composed.Bounds().Dx()))
		Expect(i.Bounds().Dy()).To(Equal(composed.Bounds().Dy()))
		for x := 0; x < i.Bounds().Dx(); x++ {
			for y := 0; y < i.Bounds().Dy(); y++ {
				Expect(i.At(x, y)).To(Equal(composed.At(x, y)))
			}
		}
	}

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

			ExpectToEqualTestImage(i, Composed)
		})
	})

	Describe("TopHeavyLayout", func() {
		var layout = TopHeavyLayout()
		var images []image.Image

		Context("with 1 image", func() {
			BeforeEach(func() {
				images = []image.Image{GirlBeforeAMirror.read()}
			})

			It("draws the composed image", func() {
				i := layout.Compose(images).Draw(400, 600)
				ExpectToEqualTestImage(i, TopHeavy1)
			})
		})

		Context("with 2 images", func() {
			BeforeEach(func() {
				images = []image.Image{
					GirlBeforeAMirror.read(),
					OldGuitarist.read(),
				}
			})

			It("draws the composed image", func() {
				i := layout.Compose(images).Draw(400, 600)
				ExpectToEqualTestImage(i, TopHeavy2)
			})
		})

		Context("with 3 images", func() {
			BeforeEach(func() {
				images = []image.Image{
					GirlBeforeAMirror.read(),
					OldGuitarist.read(),
					WomenOfAlgiers.read(),
				}
			})

			It("draws the composed image", func() {
				i := layout.Compose(images).Draw(400, 600)
				ExpectToEqualTestImage(i, TopHeavy3)
			})
		})

		Context("with 4 images", func() {
			BeforeEach(func() {
				images = []image.Image{
					GirlBeforeAMirror.read(),
					OldGuitarist.read(),
					WomenOfAlgiers.read(),
					Bullfight.read(),
				}
			})

			It("draws the composed image", func() {
				i := layout.Compose(images).Draw(400, 600)
				ExpectToEqualTestImage(i, TopHeavy4)
			})
		})
	})
})
