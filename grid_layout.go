package picasso

import "image"

type orientation bool

const (
	vertical   orientation = orientation(true)
	horizontal orientation = orientation(false)
)

func getImageOrientation(image image.Image) orientation {
	rect := image.Bounds()
	if rect.Dx() >= rect.Dy() {
		return horizontal
	} else {
		return vertical
	}
}

type gridLayout struct{}

func (l gridLayout) x() {
}

func (l gridLayout) findMainOrientation(images []image.Image) orientation {
	horizontalCount, _ := l.countMainOrientation(0, 0, images)
	if horizontalCount == 1 {
		return horizontal
	} else {
		return vertical
	}
}

// countMainOrientation recursively traverses the list of provided images and looks at what would the orienation of the composed image
// be if all of those images were put into a grid layout. It works by counting 2 horizontal (landscape) images as a single vertical
// (portrait) image and vice versa. It might ignore the last image in the list, if adding that would mess up the layout (for example, if
// the current calculated orientation is horizontal and the last picture is vertical, then the last picture will be ignored).
func (l gridLayout) countMainOrientation(horizontalCount, verticalCount int, images []image.Image) (int, int) {
	if len(images) == 0 {
		return horizontalCount, verticalCount
	}
	newHorizontalCount, newVerticalCount := l.addImageToMainOrientationCount(horizontalCount, verticalCount, images[0])
	// Ignore the last image if it would leave us in an incompatible state
	if len(images) == 1 && newHorizontalCount == 1 && newVerticalCount == 1 {
		return horizontalCount, verticalCount
	}
	return l.countMainOrientation(newHorizontalCount, newVerticalCount, images[1:])
}

func (l gridLayout) addImageToMainOrientationCount(horizontalCount, verticalCount int, image image.Image) (int, int) {
	orietation := getImageOrientation(image)
	if orietation == vertical {
		return l.addVertical(horizontalCount, verticalCount)
	} else {
		return l.addHorizontal(horizontalCount, verticalCount)
	}
}
func (l gridLayout) addHorizontal(horizontalCount, verticalCount int) (int, int) {
	if horizontalCount == 1 {
		return l.merge2Horizontals(horizontalCount+1, verticalCount)
	}
	return horizontalCount + 1, verticalCount
}

func (l gridLayout) merge2Horizontals(horizontalCount, verticalCount int) (int, int) {
	return l.addVertical(horizontalCount-2, verticalCount)
}

func (l gridLayout) addVertical(horizontalCount, verticalCount int) (int, int) {
	if verticalCount == 1 {
		return l.merge2Verticals(horizontalCount, verticalCount+1)
	}
	return horizontalCount, verticalCount + 1
}

func (l gridLayout) merge2Verticals(horizontalCount, verticalCount int) (int, int) {
	return l.addHorizontal(horizontalCount, verticalCount-2)
}
