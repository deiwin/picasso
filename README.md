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

*Picasso* also supports different automatic layouts, so that the following code:

```go
images = []image.Image{
	girlBeforeAMirror,
	oldGuitarist,
	womenOfAlgiers,
	bullfight,
	weepingWoman,
	laReve,
}
layout := picasso.GoldenSpiralLayout()
image := layout.Compose(images).Draw(600, 600)
```

Will compose an image using the golden ratio:

![automatic](https://raw.githubusercontent.com/deiwin/picasso/master/test_images/golden_spiral-6.png)

*See tests for more examples*
