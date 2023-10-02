package emboss

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/png"
	"io"
	_ "log"
	"net/url"
	"os"
	"path/filepath"

	emboss_grpc "github.com/sfomuseum/go-image-emboss/grpc"
	"google.golang.org/grpc"
)

type GrpcEmbosser struct {
	conn   *grpc.ClientConn
	client emboss_grpc.EmbosserClient
}

func init() {
	ctx := context.Background()
	RegisterEmbosser(ctx, "grpc", NewGrpcEmbosser)
}

func NewGrpcEmbosser(ctx context.Context, uri string) (Embosser, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	addr := u.Host

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		return nil, fmt.Errorf("Failed to dial '%s', %w", addr, err)
	}

	// defer conn.Close()

	client := emboss_grpc.NewEmbosserClient(conn)

	e := &GrpcEmbosser{
		conn:   conn,
		client: client,
	}

	return e, nil
}

func (e *GrpcEmbosser) EmbossImage(ctx context.Context, path string) ([]image.Image, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	im_r, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s for reading, %w", path, err)
	}

	defer im_r.Close()

	return e.EmbossImageWithReader(ctx, path, im_r)
}

func (e *GrpcEmbosser) EmbossImageWithReader(ctx context.Context, path string, im_r io.Reader) ([]image.Image, error) {

	fname := filepath.Base(path)

	body, err := io.ReadAll(im_r)

	if err != nil {
		return nil, fmt.Errorf("Failed to read %s, %w", path, err)
	}

	req := &emboss_grpc.EmbossImageRequest{
		Filename: fname,
		Body:     body,
	}

	rsp, err := e.client.EmbossImage(ctx, req)

	if err != nil {
		return nil, fmt.Errorf("Failed to emboss image, %w", err)
	}

	images := make([]image.Image, len(rsp.Body))

	for idx, data := range rsp.Body {

		r := bytes.NewReader(data)
		im, _, err := image.Decode(r)

		if err != nil {
			return nil, fmt.Errorf("Failed to decode image at offset %d, %w", idx, err)
		}

		images[idx] = im
	}

	return images, nil
}

func (e *GrpcEmbosser) Close(ctx context.Context) error {
	return e.conn.Close()
}
