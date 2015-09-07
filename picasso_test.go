package picasso_test

import (
	"image"
	"image/color"
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

	PictureWithBorder             TestImage = "./test_images/picture_with_border.png"
	VerticalSplitWithBorder       TestImage = "./test_images/vertical_split_with_border.png"
	VerticalSplitWithThinBorder   TestImage = "./test_images/vertical_split_with_thin_border.png"
	HorizontalSplitWithBorder     TestImage = "./test_images/horizontal_split_with_border.png"
	HorizontalSplitWithThinBorder TestImage = "./test_images/horizontal_split_with_thin_border.png"

	Composed TestImage = "./test_images/composed.png"

	TopHeavy1 TestImage = "./test_images/top_heavy-1.png"
	TopHeavy2 TestImage = "./test_images/top_heavy-2.png"
	TopHeavy3 TestImage = "./test_images/top_heavy-3.png"
	TopHeavy4 TestImage = "./test_images/top_heavy-4.png"

	GoldenSpiral1          TestImage = "./test_images/golden_spiral-1.png"
	GoldenSpiral2          TestImage = "./test_images/golden_spiral-2.png"
	GoldenSpiral3          TestImage = "./test_images/golden_spiral-3.png"
	GoldenSpiral4          TestImage = "./test_images/golden_spiral-4.png"
	GoldenSpiral5          TestImage = "./test_images/golden_spiral-5.png"
	GoldenSpiral6          TestImage = "./test_images/golden_spiral-6.png"
	GoldenSpiralWithBorder TestImage = "./test_images/golden_spiral_with_border.png"
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

	Describe("Picture", func() {
		Describe("DrawWithBorder", func() {
			It("adds the borders", func() {
				i := Picture{Bullfight.read()}.DrawWithBorder(400, 200, color.RGBA{0xff, 0x00, 0x00, 0xff}, 2)
				ExpectToEqualTestImage(i, PictureWithBorder)
			})
		})
	})

	Describe("VerticalSplit", func() {
		Describe("DrawWithBorder", func() {
			It("adds the borders", func() {
				i := VerticalSplit{
					Ratio: 1,
					Left:  Picture{OldGuitarist.read()},
					Right: Picture{WomenOfAlgiers.read()},
				}.DrawWithBorder(400, 400, color.RGBA{0xff, 0x00, 0x00, 0xff}, 2)
				ExpectToEqualTestImage(i, VerticalSplitWithBorder)
			})

			It("adds thin borders", func() {
				i := VerticalSplit{
					Ratio: 1,
					Left:  Picture{OldGuitarist.read()},
					Right: Picture{WomenOfAlgiers.read()},
				}.DrawWithBorder(400, 400, color.RGBA{0xff, 0x00, 0x00, 0xff}, 1)
				ExpectToEqualTestImage(i, VerticalSplitWithThinBorder)
			})
		})
	})

	Describe("HorizontalSplit", func() {
		Describe("DrawWithBorder", func() {
			It("adds the borders", func() {
				i := HorizontalSplit{
					Ratio:  1,
					Top:    Picture{Bullfight.read()},
					Bottom: Picture{WomenOfAlgiers.read()},
				}.DrawWithBorder(400, 400, color.RGBA{0xff, 0x00, 0x00, 0xff}, 2)
				ExpectToEqualTestImage(i, HorizontalSplitWithBorder)
			})

			It("adds thin borders", func() {
				i := HorizontalSplit{
					Ratio:  1,
					Top:    Picture{Bullfight.read()},
					Bottom: Picture{WomenOfAlgiers.read()},
				}.DrawWithBorder(400, 400, color.RGBA{0xff, 0x00, 0x00, 0xff}, 1)
				ExpectToEqualTestImage(i, HorizontalSplitWithThinBorder)
			})
		})
	})

	Describe("Node", func() {
		Describe("Draw", func() {
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

			Describe("with borders", func() {
				It("draws the composed image with borders", func() {
					i := layout.Compose(images).DrawWithBorder(600, 600, color.RGBA{0xaf, 0xaf, 0xaf, 0xff}, 2)
					ExpectToEqualTestImage(i, GoldenSpiralWithBorder)
				})
			})
		})
	})
})
