package emboss

import (
	"bytes"
	"context"
	"crypto/tls"
	_ "crypto/x509"
	"fmt"
	"image"
	_ "image/png"
	"io"
	_ "io/ioutil"
	_ "log"
	"net/url"
	"os"
	"path/filepath"

	emboss_grpc "github.com/sfomuseum/go-image-emboss/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	q := u.Query()

	tls_cert := q.Get("tls-certificate")
	tls_key := q.Get("tls-key")

	addr := u.Host

	var opts []grpc.DialOption
	opts = append(opts)

	if tls_cert != "" && tls_key != "" {

		/*
			ca_cert, err := ioutil.ReadFile(tls_cert)

			if err != nil {
				return nil, fmt.Errorf("Failed to create CA certificate, %w", err)
			}

			cert_pool := x509.NewCertPool()

			ok := cert_pool.AppendCertsFromPEM(ca_cert)

			if !ok {
				return nil, fmt.Errorf("Failed to append CA certificate, %w", err)
			}
		*/

		cert, err := tls.LoadX509KeyPair(tls_cert, tls_key)

		if err != nil {
			return nil, fmt.Errorf("Failed to load TLS pair, %w", err)
		}

		tls_config := &tls.Config{
			Certificates: []tls.Certificate{cert},
			// RootCAs:      cert_pool,
			InsecureSkipVerify: true,
		}

		tls_credentials := credentials.NewTLS(tls_config)
		opts = append(opts, grpc.WithTransportCredentials(tls_credentials))

	} else {
		opts = append(opts, grpc.WithInsecure())
	}

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

func (e *GrpcEmbosser) EmbossImage(ctx context.Context, path string, combined bool) ([]image.Image, error) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	im_r, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s for reading, %w", path, err)
	}

	defer im_r.Close()

	return e.EmbossImageWithReader(ctx, im_r, path, combined)
}

func (e *GrpcEmbosser) EmbossImageWithReader(ctx context.Context, im_r io.Reader, path string, combined bool) ([]image.Image, error) {

	fname := filepath.Base(path)

	body, err := io.ReadAll(im_r)

	if err != nil {
		return nil, fmt.Errorf("Failed to read %s, %w", path, err)
	}

	req := &emboss_grpc.EmbossImageRequest{
		Filename: fname,
		Body:     body,
		Combined: combined,
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
