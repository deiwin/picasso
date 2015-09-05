package picasso

import (
	"image"
	"image/draw"

	"github.com/disintegration/gift"
)

type Node interface {
	Draw(width, height int) image.Image
}

type Picture struct {
	Picture image.Image
}

func (n Picture) Draw(width, height int) image.Image {
	g := gift.New(
		gift.ResizeToFill(width, height, gift.LanczosResampling, gift.CenterAnchor),
	)
	dst := image.NewRGBA(g.Bounds(n.Picture.Bounds()))
	g.Draw(dst, n.Picture)
	return dst
}

type VerticalSplit struct {
	Left  Node
	Right Node
	Ratio float32
}

func (n VerticalSplit) Draw(width, height int) image.Image {
	rightWidth := n.rightWidth(width)
	leftWidth := width - rightWidth

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	leftImage := n.Left.Draw(leftWidth, height)
	leftRect := image.Rect(0, 0, leftWidth, height)
	draw.Draw(dst, leftRect, leftImage, image.ZP, draw.Over)

	rightImage := n.Right.Draw(rightWidth, height)
	rightRect := image.Rect(leftWidth, 0, width, height)
	draw.Draw(dst, rightRect, rightImage, image.ZP, draw.Over)

	return dst
}

func (n VerticalSplit) rightWidth(width int) int {
	// Go doesn't have a simple round function and the rounding direction doesn't really matter here,
	// so we'll just coerce the result to an int which discards the fraction.
	return int(float32(width) / (n.Ratio + 1))
}
