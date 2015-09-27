package picasso

import "image"

type orientation bool

const (
	horizontal orientation = orientation(false)
	vertical   orientation = orientation(true)
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

func (l gridLayout) doThings(images []image.Image) {
	orientation, imagesToBeComposed := l.composableSubset(images)
	if orientation == horizontal {
		return l.splitVertically(imagesToBeComposed)
	} else {
		return l.splitHorizontally(imagesToBeComposed)
	}
}

// composableSubset returns either all or all but the last image provided depending on if all images can be used to create the layout
// without any gaps. It also returns what orientation the resulting images would take when composed.
func (l gridLayout) composableSubset(images []image.Image) (orientation, []image.Image) {
	horizontalCount, verticalCount := l.countComposedOrientation(0, 0, images)
	if horizontalCount == 1 && verticalCount == 1 {
		lastImageOrientation := getImageOrientation(images[len(images)-1])
		if lastImageOrientation == horizontal {
			return vertical, images[:len(images)-1]
		} else {
			return horizontal, images[:len(images)-1]
		}
	}

	if horizontalCount == 1 {
		return horizontal, images
	} else {
		return vertical, images
	}
}

// countComposedOrientation recursively traverses the list of provided images and looks at what would the orienation of the composed image
// be if all of those images were put into a grid layout. It works by counting 2 horizontal (landscape) images as a single vertical
// (portrait) image and vice versa.
func (l gridLayout) countComposedOrientation(horizontalCount, verticalCount int, images []image.Image) (int, int) {
	if len(images) == 0 {
		return horizontalCount, verticalCount
	}
	newHorizontalCount, newVerticalCount := l.addImageToComposedOrientationCount(horizontalCount, verticalCount, images[0])
	return l.countComposedOrientation(newHorizontalCount, newVerticalCount, images[1:])
}

func (l gridLayout) addImageToComposedOrientationCount(horizontalCount, verticalCount int, image image.Image) (int, int) {
	orientation := getImageOrientation(image)
	if orientation == vertical {
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
