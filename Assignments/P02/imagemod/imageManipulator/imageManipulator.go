// imagemod/imageManipulator/imageManipulator.go

package imageManipulator

// This will prompt you to go get the package :-)
import (
	"github.com/fogleman/gg"
)

// ImageManipulator represents an image manipulation tool.
type ImageManipulator struct {
	Image     *gg.Context
	ImagePath string
}

// NewImageManipulator creates a new ImageManipulator instance.
func NewImageManipulator(width, height int) *ImageManipulator {
	img := gg.NewContext(width, height)
	return &ImageManipulator{Image: img}
}

// SaveToFile saves the manipulated image to a file.
func (im *ImageManipulator) SaveToFile(filename string) error {
	return im.Image.SavePNG(filename)
}

// DrawRectangle draws a rectangle on the image.
func (im *ImageManipulator) DrawRectangle(x, y, width, height float64) {
	im.Image.DrawRectangle(x, y, width, height)
	im.Image.Stroke()
}

// NewImageManipulatorWithImage loads an existing image and creates an ImageManipulator instance.
func NewImageManipulatorWithImage(imagePath string) (*ImageManipulator, error) {
	img, err := gg.LoadImage(imagePath)
	// Error handling in case the image fails to load.....which it hopefully won't.
	if err != nil {
		return nil, err
	}
	// Copies an image to a new context that will be returned
	ctx := gg.NewContextForImage(img)
	// Returns the image and its path
	return &ImageManipulator{Image: ctx, ImagePath: imagePath}, nil
}
