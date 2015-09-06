package picasso

import (
	"image"
	"math"
)

type Layout interface {
	Compose([]image.Image) Node
}

// TopHeavyLayout creates a layout that works well for up to 4 images with a landscape aspect ratio.
// In this layout, 2 images are shown with equal sizes one atop the other. With 3 images the bottom
// part will get split in half and 2 images will be fit into that part side by side. With 4 and more
// images, the first image will be on top and rest of the images will be on a single row below, all
// with equal widths. Also, with 4 and more images, the top part's height will be 2 times that of the
// bottom one.
func TopHeavyLayout() Layout {
	return topHeavy{}
}

type topHeavy struct{}

func (l topHeavy) Compose(images []image.Image) Node {
	if len(images) == 0 {
		return nil
	} else if len(images) == 1 {
		return Picture{images[0]}
	} else if len(images) == 2 {
		return HorizontalSplit{
			Ratio:  1,
			Top:    Picture{images[0]},
			Bottom: Picture{images[1]},
		}
	} else if len(images) == 3 {
		return HorizontalSplit{
			Ratio: 1,
			Top:   Picture{images[0]},
			Bottom: VerticalSplit{
				Ratio: 1,
				Left:  Picture{images[1]},
				Right: Picture{images[2]},
			},
		}
	}
	return HorizontalSplit{
		Ratio:  2,
		Top:    Picture{images[0]},
		Bottom: createJustifiedVerticalSplits(images[1:]),
	}
}

func createJustifiedVerticalSplits(images []image.Image) Node {
	if len(images) == 0 {
		return nil
	} else if len(images) == 1 {
		return Picture{images[0]}
	}
	tail := images[1:]
	return VerticalSplit{
		Ratio: float32(1) / float32(len(tail)),
		Left:  Picture{images[0]},
		Right: createJustifiedVerticalSplits(tail),
	}
}

// GoldenSpiralLayout will create a layout that creates splits following a golden spiral that
// starts moving to the right and to the bottom (the most common version).
func GoldenSpiralLayout() Layout {
	return goldenSpiral{}
}

type goldenSpiral struct{}

func (l goldenSpiral) Compose(images []image.Image) Node {
	return l.splitRight(images)
}

func (l goldenSpiral) splitRight(images []image.Image) Node {
	if len(images) == 0 {
		return nil
	} else if len(images) == 1 {
		return Picture{images[0]}
	}

	tail := images[1:]
	return VerticalSplit{
		Ratio: math.Phi,
		Left:  Picture{images[0]},
		Right: l.splitBelow(tail),
	}
}

func (l goldenSpiral) splitBelow(images []image.Image) Node {
	if len(images) == 0 {
		return nil
	} else if len(images) == 1 {
		return Picture{images[0]}
	}

	tail := images[1:]
	return HorizontalSplit{
		Ratio:  math.Phi,
		Top:    Picture{images[0]},
		Bottom: l.splitLeft(tail),
	}
}

func (l goldenSpiral) splitLeft(images []image.Image) Node {
	if len(images) == 0 {
		return nil
	} else if len(images) == 1 {
		return Picture{images[0]}
	}

	tail := images[1:]
	return VerticalSplit{
		Ratio: float32(1) / math.Phi,
		Left:  l.splitAbove(tail),
		Right: Picture{images[0]},
	}
}

func (l goldenSpiral) splitAbove(images []image.Image) Node {
	if len(images) == 0 {
		return nil
	} else if len(images) == 1 {
		return Picture{images[0]}
	}

	tail := images[1:]
	return HorizontalSplit{
		Ratio:  float32(1) / math.Phi,
		Top:    l.splitRight(tail),
		Bottom: Picture{images[0]},
	}
}
