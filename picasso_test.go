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
	WeepingWoman      TestImage = "./test_images/picasso-the_weeping_woman.jpg"
	LaReve            TestImage = "./test_images/picasso-la_reve.jpg"

	Composed TestImage = "./test_images/composed.png"

	TopHeavy1 TestImage = "./test_images/top_heavy-1.png"
	TopHeavy2 TestImage = "./test_images/top_heavy-2.png"
	TopHeavy3 TestImage = "./test_images/top_heavy-3.png"
	TopHeavy4 TestImage = "./test_images/top_heavy-4.png"

	GoldenSpiral1 TestImage = "./test_images/golden_spiral-1.png"
	GoldenSpiral2 TestImage = "./test_images/golden_spiral-2.png"
	GoldenSpiral3 TestImage = "./test_images/golden_spiral-3.png"
	GoldenSpiral4 TestImage = "./test_images/golden_spiral-4.png"
	GoldenSpiral5 TestImage = "./test_images/golden_spiral-5.png"
	GoldenSpiral6 TestImage = "./test_images/golden_spiral-6.png"
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
		expected := testImage.read()
		Expect(i.Bounds().Dx()).To(Equal(expected.Bounds().Dx()))
		Expect(i.Bounds().Dy()).To(Equal(expected.Bounds().Dy()))
		for x := 0; x < i.Bounds().Dx(); x++ {
			for y := 0; y < i.Bounds().Dy(); y++ {
				expectedColor := expected.At(x, y)
				er, eg, eb, ea := expectedColor.RGBA()
				r, g, b, a := i.At(x, y).RGBA()
				// There might be some differences while creating the image on different machines.
				// To account for that, expect the images to be similar, but not exactly the same.
				Expect(r).To(BeNumerically("~", er, 2))
				Expect(g).To(BeNumerically("~", eg, 2))
				Expect(b).To(BeNumerically("~", eb, 2))
				Expect(a).To(BeNumerically("~", ea, 2))
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
			It("returns nil", func() {
				l := layout.Compose(images)
				Expect(l).To(BeNil())
			})
		})

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

	Describe("GoldenSpiralLayout", func() {
		var layout = GoldenSpiralLayout()
		var images []image.Image

		Context("with 1 image", func() {
			It("returns nil", func() {
				l := layout.Compose(images)
				Expect(l).To(BeNil())
			})
		})

		Context("with 1 image", func() {
			BeforeEach(func() {
				images = []image.Image{GirlBeforeAMirror.read()}
			})

			It("draws the composed image", func() {
				i := layout.Compose(images).Draw(600, 600)
				ExpectToEqualTestImage(i, GoldenSpiral1)
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
				i := layout.Compose(images).Draw(600, 600)
				ExpectToEqualTestImage(i, GoldenSpiral2)
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
				i := layout.Compose(images).Draw(600, 600)
				ExpectToEqualTestImage(i, GoldenSpiral3)
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
				i := layout.Compose(images).Draw(600, 600)
				ExpectToEqualTestImage(i, GoldenSpiral4)
			})
		})

		Context("with 5 images", func() {
			BeforeEach(func() {
				images = []image.Image{
					GirlBeforeAMirror.read(),
					OldGuitarist.read(),
					WomenOfAlgiers.read(),
					Bullfight.read(),
					WeepingWoman.read(),
				}
			})

			It("draws the composed image", func() {
				i := layout.Compose(images).Draw(600, 600)
				ExpectToEqualTestImage(i, GoldenSpiral5)
			})
		})

		Context("with 6 images", func() {
			BeforeEach(func() {
				images = []image.Image{
					GirlBeforeAMirror.read(),
					OldGuitarist.read(),
					WomenOfAlgiers.read(),
					Bullfight.read(),
					WeepingWoman.read(),
					LaReve.read(),
				}
			})

			It("draws the composed image", func() {
				i := layout.Compose(images).Draw(600, 600)
				ExpectToEqualTestImage(i, GoldenSpiral6)
			})
		})
	})
})
