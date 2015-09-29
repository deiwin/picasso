# Picasso, a Go image composer

[![Build Status](https://travis-ci.org/deiwin/picasso.svg?branch=master)](https://travis-ci.org/deiwin/picasso)
[![Coverage](http://gocover.io/_badge/github.com/deiwin/picasso?0)](http://gocover.io/github.com/deiwin/picasso)
[![GoDoc](https://godoc.org/github.com/deiwin/picasso?status.svg)](https://godoc.org/github.com/deiwin/picasso)

## Example
### Manual layout handling

The following code:

```go
image := picasso.HorizontalSplit{
	Ratio: 2,
	Top:   picasso.Picture{bullfight},
	Bottom: picasso.VerticalSplit{
		Ratio: 0.5,
		Left:  picasso.Picture{girlBeforeAMirror},
		Right: picasso.VerticalSplit{
			Ratio: 1,
			Left:  picasso.Picture{oldGuitarist},
			Right: picasso.Picture{womenOfAlgiers},
		},
	},
}.Draw(400, 600)
```

Will compose the following image:

![manual](https://raw.githubusercontent.com/deiwin/picasso/master/test_images/composed.png)

### Automatic layouts

*Picasso* also supports different automatic layouts and borders, so that the following code:

```go
images := []image.Image{
	girlBeforeAMirror,
	oldGuitarist,
	womenOfAlgiers,
	bullfight,
	weepingWoman,
	laReve,
}
layout := picasso.GoldenSpiralLayout()
gray := color.RGBA{0xaf, 0xaf, 0xaf, 0xff}
image := layout.Compose(images).DrawWithBorder(600, 600, gray, 2)
```

Will compose an image using the golden ratio:

![automatic](https://raw.githubusercontent.com/deiwin/picasso/master/test_images/golden_spiral_with_border.png)

Or one could use the GridLayout:
```go
images := []image.Image{...}
gray := color.RGBA{0xaf, 0xaf, 0xaf, 0xff}
image := picasso.DrawGridLayoutWithBorder(images, 800, gray, 2)
```
to compose larger sets of images:

![composed](https://cloud.githubusercontent.com/assets/2261897/10125748/c22d5144-6588-11e5-8962-8458313ff0bf.jpg)

*See tests for more examples*
