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

func (l gridLayout) doThings(images []image.Image) Node {
	orientation, imagesToBeComposed := l.composableSubset(images)
	if orientation == horizontal {
		return l.splitVertically(imagesToBeComposed)
	} else {
		return l.splitHorizontally(imagesToBeComposed)
	}
}

func (l gridLayout) splitVertically(images []image.Image) Node {
	if len(images) == 1 {
		return Picture{images[0]}
	}
	// add + 1 to effectively round the division up as we want initially there to be more images on the left side
	// so that when we start moving images from one side to the other, the numbers would stay more or less in balance
	midPoint := (len(images) + 1) / 2
	proposedLeftImages := images[:midPoint]
	proposedRightImages := images[midPoint:]
	proposedLeftHorizontalCount, proposedLeftVerticalCount := l.countComposedOrientation(0, 0, proposedLeftImages)
	proposedRightHorizontalCount, proposedRightVerticalCount := l.countComposedOrientation(0, 0, proposedRightImages)
	// now we need to redistribute the images so that both sides would compose to vertical orientation
	// NB: this method expects this to be possible
	if proposedLeftHorizontalCount == 1 && proposedLeftVerticalCount == 0 && proposedRightHorizontalCount == 1 && proposedRightVerticalCount == 1 {
		leftHasHorizontal, lastLeftHorizontalIndex := indexOfLastHorizontalImage(proposedLeftImages)
		rightHasVertical, lastRightVerticalIndex := indexOfLastVerticalImage(proposedRightImages)
		if leftHasHorizontal && rightHasVertical {
			leftImages, rightImages := swapImage(proposedLeftImages, lastLeftHorizontalIndex, proposedRightImages, lastRightVerticalIndex)
			return VerticalSplit{
				Ratio: 1,
				Left:  l.splitHorizontally(leftImages),
				Right: l.splitHorizontally(rightImages),
			}
		}
		leftHasVertical, lastLeftVerticalIndex := indexOfLastVerticalImage(proposedLeftImages)
		if leftHasVertical {
			leftImages, rightImages := move1ImageOver(proposedLeftImages, lastLeftVerticalIndex, proposedRightImages)
			return VerticalSplit{
				Ratio: 1,
				Left:  l.splitHorizontally(leftImages),
				Right: l.splitHorizontally(rightImages),
			}
		}
		nextToLastLeftHorizontalIndex := indexOfNextToLastHorizontalImage(proposedLeftImages)
		leftImages, rightImages := move2ImagesOver(proposedLeftImages, nextToLastLeftHorizontalIndex, lastLeftHorizontalIndex, proposedRightImages)
		return VerticalSplit{
			Ratio: 1,
			Left:  l.splitHorizontally(leftImages),
			Right: l.splitHorizontally(rightImages),
		}
	} else if proposedLeftHorizontalCount == 1 && proposedLeftVerticalCount == 1 && proposedRightHorizontalCount == 1 && proposedRightVerticalCount == 0 {

	}
	return VerticalSplit{
		Ratio: 1,
		Left:  l.splitHorizontally(proposedLeftImages),
		Right: l.splitHorizontally(proposedRightImages),
	}
}

func (l gridLayout) splitHorizontally(images []image.Image) Node {
	if len(images) == 1 {
		return Picture{images[0]}
	}
	return nil
}

func move1ImageOver(aImages []image.Image, aIndex int, bImages []image.Image) ([]image.Image, []image.Image) {
	newAImages := make([]image.Image, len(aImages)-1)
	copy(newAImages[:aIndex], aImages[:aIndex])
	copy(newAImages[aIndex:], aImages[aIndex+1:])
	newBImages := append(bImages, aImages[aIndex])
	return newAImages, newBImages
}

// move2ImagesOver expects aIndex1 < aIndex2
func move2ImagesOver(aImages []image.Image, aIndex1, aIndex2 int, bImages []image.Image) ([]image.Image, []image.Image) {
	newAImages := make([]image.Image, len(aImages)-2)
	copy(newAImages[:aIndex1], aImages[:aIndex1])
	copy(newAImages[aIndex1:aIndex2-1], aImages[aIndex1+1:aIndex2])
	copy(newAImages[aIndex2-1:], aImages[aIndex2+1:])
	newBImages := append(bImages, aImages[aIndex1], aImages[aIndex2])
	return newAImages, newBImages
}

func swapImage(aImages []image.Image, aIndex int, bImages []image.Image, bIndex int) ([]image.Image, []image.Image) {
	newAImages := make([]image.Image, len(aImages))
	newBImages := make([]image.Image, len(bImages))
	for i, image := range aImages {
		if i == aIndex {
			newAImages[i] = bImages[bIndex]
			continue
		}
		newAImages[i] = image
	}
	for i, image := range bImages {
		if i == bIndex {
			newBImages[i] = aImages[aIndex]
			continue
		}
		newBImages[i] = image
	}
	return newAImages, newBImages
}

func indexOfLastVerticalImage(images []image.Image) (hasVerticalImage bool, index int) {
	for i := len(images) - 1; i >= 0; i-- {
		orientation := getImageOrientation(images[i])
		if orientation == vertical {
			return true, i
		}
	}
	return false, -1
}

func indexOfLastHorizontalImage(images []image.Image) (hasHorizontalImage bool, index int) {
	for i := len(images) - 1; i >= 0; i-- {
		orientation := getImageOrientation(images[i])
		if orientation == horizontal {
			return true, i
		}
	}
	return false, -1
}

func indexOfNextToLastHorizontalImage(images []image.Image) int {
	foundLast := false
	for i := len(images) - 1; i >= 0; i-- {
		orientation := getImageOrientation(images[i])
		if orientation == horizontal {
			if !foundLast {
				foundLast = true
				continue
			}
			return i
		}
	}
	return -1
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
