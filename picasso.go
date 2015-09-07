package picasso

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/disintegration/gift"
)

type Node interface {
	Draw(width, height int) image.Image
	DrawWithBorder(width, height int, borderColor color.Color, borderWidth int) image.Image
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

func (n Picture) DrawWithBorder(width, height int, borderColor color.Color, borderWidth int) image.Image {
	fullRect := image.Rect(0, 0, width, height)
	inBorderRect := image.Rect(borderWidth, borderWidth, width-borderWidth, height-borderWidth)

	dst := image.NewRGBA(fullRect)

	background := image.NewUniform(borderColor)
	draw.Draw(dst, fullRect, background, image.ZP, draw.Over)

	inBorderPicture := n.Draw(inBorderRect.Dx(), inBorderRect.Dy())
	draw.Draw(dst, inBorderRect, inBorderPicture, image.ZP, draw.Over)

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

func (n VerticalSplit) DrawWithBorder(width, height int, borderColor color.Color, borderWidth int) image.Image {
	// + borderWidth, because we basically draw both sides with their full borders, but then make the right border of
	// the left image and the left border of the right image overlap
	rightWithBorderWidth := n.rightWidth(width + borderWidth)
	leftWithBorderWidth := (width + borderWidth) - rightWithBorderWidth
	fullRect := image.Rect(0, 0, width, height)
	leftWithBorderRect := image.Rect(0, 0, leftWithBorderWidth, height)
	rightWithBorderRect := image.Rect(leftWithBorderWidth-borderWidth, 0, width, height)

	dst := image.NewRGBA(fullRect)

	leftWithBorderImage := n.Left.DrawWithBorder(leftWithBorderRect.Dx(), leftWithBorderRect.Dy(), borderColor, borderWidth)
	draw.Draw(dst, leftWithBorderRect, leftWithBorderImage, image.ZP, draw.Over)

	rightWithBorderImage := n.Right.DrawWithBorder(rightWithBorderRect.Dx(), rightWithBorderRect.Dy(), borderColor, borderWidth)
	draw.Draw(dst, rightWithBorderRect, rightWithBorderImage, image.ZP, draw.Over)

	return dst
}

func (n VerticalSplit) rightWidth(width int) int {
	// Go doesn't have a simple round function and the rounding direction doesn't really matter here,
	// so we'll just coerce the result to an int which discards the fraction.
	return int(float32(width) / (n.Ratio + 1))
}

type HorizontalSplit struct {
	Top    Node
	Bottom Node
	Ratio  float32
}

func (n HorizontalSplit) Draw(width, height int) image.Image {
	bottomHeight := n.bottomHeight(height)
	topHeight := height - bottomHeight

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	topImage := n.Top.Draw(width, topHeight)
	topRect := image.Rect(0, 0, width, topHeight)
	draw.Draw(dst, topRect, topImage, image.ZP, draw.Over)

	bottomImage := n.Bottom.Draw(width, bottomHeight)
	bottomRect := image.Rect(0, topHeight, width, height)
	draw.Draw(dst, bottomRect, bottomImage, image.ZP, draw.Over)

	return dst
}

func (n HorizontalSplit) DrawWithBorder(width, height int, borderColor color.Color, borderWidth int) image.Image {
	// + borderWidth, because we basically draw both sides with their full borders, but then make the bottom border of
	// the top image and the top border of the bottom image overlap
	bottomWithBorderHeight := n.bottomHeight(height + borderWidth)
	topWithBorderHeight := (height + borderWidth) - bottomWithBorderHeight
	fullRect := image.Rect(0, 0, width, height)
	topWithBorderRect := image.Rect(0, 0, width, topWithBorderHeight)
	bottomWithBorderRect := image.Rect(0, topWithBorderHeight-borderWidth, width, height)

	dst := image.NewRGBA(fullRect)

	topWithBorderImage := n.Top.DrawWithBorder(topWithBorderRect.Dx(), topWithBorderRect.Dy(), borderColor, borderWidth)
	draw.Draw(dst, topWithBorderRect, topWithBorderImage, image.ZP, draw.Over)

	bottomWithBorderImage := n.Bottom.DrawWithBorder(bottomWithBorderRect.Dx(), bottomWithBorderRect.Dy(), borderColor, borderWidth)
	draw.Draw(dst, bottomWithBorderRect, bottomWithBorderImage, image.ZP, draw.Over)

	return dst
}

func (n HorizontalSplit) bottomHeight(height int) int {
	// Go doesn't have a simple round function and the rounding direction doesn't really matter here,
	// so we'll just coerce the result to an int which discards the fraction.
	return int(float32(height) / (n.Ratio + 1))
}
