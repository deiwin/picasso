package picasso

import "image"

type Layout interface {
	Compose([]image.Image) Node
}

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
