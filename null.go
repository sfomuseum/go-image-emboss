package emboss

import (
	"context"
	"image"
	"io"
)

type NullEmbosser struct {
}

func init() {
	ctx := context.Background()
	RegisterEmbosser(ctx, "null", NewNullEmbosser)
}

func NewNullEmbosser(ctx context.Context, uri string) (Embosser, error) {
	e := &NullEmbosser{}
	return e, nil
}

func (e *NullEmbosser) EmbossImage(ctx context.Context, path string) ([]image.Image, error) {
	return []image.Image{}, nil
}

func (e *NullEmbosser) EmbossImageWithReader(ctx context.Context, path string, r io.Reader) ([]image.Image, error) {
	return []image.Image{}, nil
}

func (e *NullEmbosser) Close(ctx context.Context) error {
	return nil
}
